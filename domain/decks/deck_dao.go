package decks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/errors"
)

func (deck *Deck) Get() *errors.RestErr {
	err := mongo_db.DB.Collection("decks").FindOne(context.TODO(), bson.D{{"ID", deck.ID}}).Decode(&deck)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return nil
}
