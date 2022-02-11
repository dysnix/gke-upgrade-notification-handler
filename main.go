package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	err := pullMsg("zksync", "gke-upgrade")
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
