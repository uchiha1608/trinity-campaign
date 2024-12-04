package models

import "time"

type User struct {
	ID                   string    `json:"id"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone"`
	RegisterAt           time.Time `json:"register_at"`
	IsFirst100           bool      `json:"is_first_100"`
	CampaignLinkUsed     bool      `json:"campaign_link_used"`
	ActiveSubscriptionID string    `json:"active_subscription_id"`
}
