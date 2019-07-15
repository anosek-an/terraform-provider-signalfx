/*
 * Teams API
 *
 *  ## Overview An API for creating, retrieving, updating, and deleting teams ## Authentication To authenticate with SignalFx, the following operations require a session  token associated with a SignalFx user that has administrative privileges:<br>   * Create a team - **POST** `/team`   * Update a team - **PUT** `/team/{id}`   * Delete a team - **DELETE** `/team/{id}`   * Update team members - **PUT** `/team/{id}/members`  You can authenticate the following operations with either an org token or a session token. The session token  doesn't need to be associated with a SignalFx user that has administrative privileges:<br>   * Retrieve teams using a query - **GET** `/team`   * Retrieve a team using its ID - **GET** `/team/{id}`
 *
 * API version: 3.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package team

import "encoding/json"

// Notification properties for an alert sent to the team
type Notification struct {
	Type  string
	Value interface{}
}

func (n *Notification) UnmarshalJSON(data []byte) error {
	var typ struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &typ); err != nil {
		return err
	}
	n.Type = typ.Type
	switch typ.Type {
	case "BigPanda":
		n.Value = &BigPandaNotification{}
	case "Email":
		n.Value = &EmailNotification{}
	case "Office365":
		n.Value = &Office365Notification{}
	case "OpsGenie":
		n.Value = &OpsgenieNotification{}
	case "PagerDuty":
		n.Value = &PagerDutyNotification{}
	case "ServiceNow":
		n.Value = &ServiceNowNotification{}
	case "Slack":
		n.Value = &SlackNotification{}
	case "Team":
		n.Value = &TeamNotification{}
	case "TeamEmail":
		n.Value = &TeamEmailNotification{}
	case "VictorOps":
		n.Value = &VictorOpsNotification{}
	case "Webhook":
		n.Value = &WebhookNotification{}
	case "XMatters":
		n.Value = &XMattersNotification{}
	}
	return json.Unmarshal(data, n.Value)
}

// MarshalJSON unwraps the `Value` and makes this marshal in the way we expect.
func (n *Notification) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Value)
}
