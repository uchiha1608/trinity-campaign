package services

import (
	"fmt"
	"math/rand"
	"time"
	"trinity-campaign/database"
	"trinity-campaign/models"
)

var vouchers = []models.Voucher{}

// GenerateVoucher generate voucher based on user ID and campaignID
func GenerateVoucher(userID, campaignID int, discountAvailable bool) (models.Voucher, error) {
	// Generate random voucher code
	rand.NewSource(time.Now().UnixNano())
	voucherCode := fmt.Sprintf("VOUCHER-%d", rand.Intn(100000))

	// Create voucher
	voucher := models.Voucher{
		UserID:            userID,
		CampaignID:        campaignID,
		Code:              voucherCode,
		Status:            "active",
		ExpiryDate:        time.Now().Add(72 * time.Hour), // 3 days expiry
		DiscountAvailable: discountAvailable,
	}
	_, err := database.InsertVoucher(voucher)
	return voucher, err
}

// ValidateVoucher validate voucher based on voucher code to check its status and expiry date
func ValidateVoucher(voucherCode string) (*models.Voucher, error) {
	for _, voucher := range vouchers {
		if voucher.Code == voucherCode && voucher.Status == "active" && voucher.ExpiryDate.After(time.Now()) {
			return &voucher, nil
		}
	}
	return nil, fmt.Errorf("invalid or expired voucher")
}
