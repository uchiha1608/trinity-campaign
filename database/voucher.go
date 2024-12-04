package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertVoucher(voucher interface{}) (*mongo.InsertOneResult, error) {
	result, err := voucherCollection.InsertOne(context.Background(), voucher)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindVoucherByCode(code string) (bson.M, error) {
	var voucher bson.M
	err := voucherCollection.FindOne(context.Background(), bson.M{"voucher_code": code}).Decode(&voucher)
	if err != nil {
		return nil, err
	}
	return voucher, nil
}
