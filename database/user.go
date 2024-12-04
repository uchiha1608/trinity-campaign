package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUser(user interface{}) (*mongo.InsertOneResult, error) {
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindUserByID(userID int) (bson.M, error) {
	var user bson.M
	err := userCollection.FindOne(context.Background(), bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
