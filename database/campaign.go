package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertCampaign(campaign interface{}) (*mongo.InsertOneResult, error) {
	result, err := campaignCollection.InsertOne(context.Background(), campaign)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindCampaignByID(id string) (bson.M, error) {
	var campaign bson.M
	err := campaignCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
