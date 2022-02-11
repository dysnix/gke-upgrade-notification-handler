package main

import (
	"fmt"
	"time"
)

type Event interface {
	GetMessageText() string
}

type UpgradeEvent struct {
	ResourceType       string    `json:"resourceType"`
	Operation          string    `json:"operation"`
	OperationStartTime time.Time `json:"operationStartTime"`
	CurrentVersion     string    `json:"currentVersion"`
	TargetVersion      string    `json:"targetVersion"`
}

func (e *UpgradeEvent) GetMessageText() string {
	return fmt.Sprintf("Upgrade has started from %s to %s on %s", e.CurrentVersion, e.TargetVersion, e.ResourceType)
}

type UpgradeAvailableEvent struct {
	Version        string `json:"version"`
	ResourceType   string `json:"resourceType"`
	ReleaseChannel string `json:"releaseChannel"`
	Resource       string `json:"resource"`
}

func (e *UpgradeAvailableEvent) GetMessageText() string {
	return fmt.Sprintf("Upgrade available to %s on %s", e.Version, e.ResourceType)
}

type SecurityBulletinEvent struct {
	ResourceTypeAffected    string   `json:"resourceTypeAffected"`
	BulletinID              string   `json:"bulletinId"`
	CveIds                  []string `json:"cveIds"`
	Severity                string   `json:"severity"`
	BulletinURI             string   `json:"bulletinUri"`
	BriefDescription        string   `json:"briefDescription"`
	AffectedSupportedMinors []string `json:"affectedSupportedMinors"`
	PatchedVersions         []string `json:"patchedVersions"`
	SuggestedUpgradeTarget  string   `json:"suggestedUpgradeTarget"`
	ManualStepsRequired     bool     `json:"manualStepsRequired"`
}

func (e *SecurityBulletinEvent) GetMessageText() string {
	return e.ResourceTypeAffected
}

func GetMessageText(event Event) string {
	return event.GetMessageText()
}
