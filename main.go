package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	projectID = os.Getenv("GCP_PROJECT_ID")
	subID     = os.Getenv("GCP_SUBSCRIPTION_ID")
)

func main() {
	if projectID == "" {
		log.Fatal("GCP_PROJECT_ID must be set")
		return
	}
	if subID == "" {
		log.Fatal("GCP_SUBSCRIPTION_ID must be set")
		return
	}
	err := pullMsg(projectID, subID)
	if err != nil {
		log.Printf("Receive: %v \n", err)
	}
}

func pullMsg(projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		var message string
		switch msg.Attributes["type_url"] {
		case "type.googleapis.com/google.container.v1beta1.UpgradeEvent":
			var e UpgradeEvent
			DecodeJson(msg.Data, &e)
			message = GetMessageText(&e)
		case "type.googleapis.com/google.container.v1beta1.UpgradeAvailableEvent":
			var e UpgradeAvailableEvent
			DecodeJson(msg.Data, &e)
			message = GetMessageText(&e)
		default:
			return
		}
		cluster := msg.Attributes["cluster_name"]
		projectId := msg.Attributes["project_id"]
		message = fmt.Sprintf("Project: %s Cluster: %s \n %s", projectId, cluster, message)
		Send(message)
		msg.Ack()
		cancel()
	})
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}
	return nil
}

func DecodeJson(data []byte, result Event) {
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Printf("DecodeJson: %v", err)
	}
}

func Send(message string) {
	mmWebHookUrl := os.Getenv("MATTERMOST_WEBHOOK_URL")
	if mmWebHookUrl != "" {
		MmSend(mmWebHookUrl, message)
	}
	log.Println(message)
}
