package services

import (
	"fmt"
	"time"
	"trinity-campaign/database"
	"trinity-campaign/models"
)

var campaigns = []models.Campaign{
	{
		ID:              1,
		Name:            "First Login Promotion",
		StartTime:       time.Now(),
		EndTime:         time.Now().AddDate(0, 0, 30), // 30 days from now
		MaxUsersAllowed: 100,
		RegisteredUsers: 50,
	},
}

// CreateCampaign create new campaign
func CreateCampaign(campaign models.Campaign) (models.Campaign, error) {
	// Logic for creating a campaign
	result, err := database.InsertCampaign(campaign)
	if err != nil {
		return campaign, err
	}
	campaign.ID = result.InsertedID.(int)
	return campaign, nil
}

// GetCampaignByID return campaign based on provided ID
func GetCampaignByID(id int) (*models.Campaign, error) {
	for _, campaign := range campaigns {
		if campaign.ID == id {
			return &campaign, nil
		}
	}
	return nil, fmt.Errorf("campaign not found")
}
