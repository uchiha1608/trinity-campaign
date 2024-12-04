package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPurchase(purchase interface{}) (*mongo.InsertOneResult, error) {
	result, err := purchaseCollection.InsertOne(context.Background(), purchase)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindPurchaseByUserID(userID int) (bson.M, error) {
	var purchase bson.M
	err := purchaseCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&purchase)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func FindPurchaseByID(id string) (bson.M, error) {
	var purchase bson.M
	err := purchaseCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&purchase)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}
