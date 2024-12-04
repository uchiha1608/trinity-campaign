package models

import "time"

type Purchase struct {
	ID                 int       `json:"id" bson:"_id"`
	UserID             int       `json:"user_id"`
	SubscriptionPlanID int       `json:"subscription_plan_id"`
	PurchaseAt         time.Time `json:"purchase_at"`
	DiscountApplied    bool      `json:"discount_applied"`
	VoucherCode        string    `json:"voucher_code"`
}
