# gke-upgrade-notification-handler

Handles the notification that a cluster is being upgraded.

##Environment Variables
| VARIABLE               | Description                       | Default value |
|------------------------|-----------------------------------|---------------|
| GCP_PROJECT_ID         | `Name of your gcp project`        | -             |
| GCP_SUBSCRIPTION_ID    | `GCP pub/sub subscription id`     | -             |
| MATTERMOST_WEBHOOK_URL | `Mattermost incoming webhook url` | -             |


##Installation


##Integrations
### Mattermost
Create a Mattermost incoming webhook and set the MATTERMOST_WEBHOOK_URL variable to the one you created.
https://jeffschering.github.io/mmdocs/upgrade/developer/webhooks-incoming.html