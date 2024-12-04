package models

import "time"

type Voucher struct {
	ID                int       `json:"id"`
	UserID            int       `json:"user_id"`
	CampaignID        int       `json:"campaign_id"`
	Code              string    `json:"code"`
	Status            string    `json:"status"`
	ExpiryDate        time.Time `json:"expiry_date"`
	DiscountAvailable bool      `json:"discount_available"`
}
