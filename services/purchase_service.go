package services

import (
	"log"
	"time"
	"trinity-campaign/database"
	"trinity-campaign/models"
)

var purchases = []models.Purchase{}

// ApplyVoucherToPurchase apply available voucher to make purchase by users
func ApplyVoucherToPurchase(userID, subscriptionPlanID int, voucherCode string) (*models.Purchase, error) {
	// Validate voucher
	voucher, err := ValidateVoucher(voucherCode)
	if err != nil {
		return nil, err
	}

	// Apply discount if voucher is valid
	purchase := models.Purchase{
		UserID:             userID,
		SubscriptionPlanID: subscriptionPlanID,
		PurchaseAt:         time.Now(),
		DiscountApplied:    voucher.DiscountAvailable,
		VoucherCode:        voucher.Code,
	}

	_, err = database.InsertPurchase(&purchase)
	if err != nil {
		log.Printf("failed to create purchase")
		return nil, err
	}
	return &purchase, nil
}
