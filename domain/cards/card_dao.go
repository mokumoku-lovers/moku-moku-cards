package cards

import (
	"context"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/errors"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func (card *Card) Get() *errors.RestErr {
	err := mongo_db.DB.Collection("cards").FindOne(context.TODO(), bson.D{{"ID", card.ID}}).Decode(&card)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return nil
}
