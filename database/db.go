package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	client             *mongo.Client
	db                 *mongo.Database
	campaignCollection *mongo.Collection
	voucherCollection  *mongo.Collection
	userCollection     *mongo.Collection
	purchaseCollection *mongo.Collection
)

func InitDB() {
	mongoUrl := "mongodb://localhost:27017"
	dbName := "trinity_campaign"

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(dbName)
	campaignCollection = db.Collection("campaigns")
	voucherCollection = db.Collection("vouchers")
	userCollection = db.Collection("users")
	purchaseCollection = db.Collection("purchases")

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
