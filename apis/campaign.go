package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"trinity-campaign/models"
	"trinity-campaign/services"
)

// Create a new campaign
func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign models.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdCampaign, err := services.CreateCampaign(campaign)
	if err != nil {
		http.Error(w, "Error creating campaign", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCampaign)
}

// Get campaign status
func GetCampaignStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	campaignID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Wrong campaign ID format", http.StatusInternalServerError)
		return
	}

	// Fetch campaign from service
	campaign, err := services.GetCampaignByID(campaignID)
	if err != nil {
		http.Error(w, "Error getting campaign", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(campaign)

}
