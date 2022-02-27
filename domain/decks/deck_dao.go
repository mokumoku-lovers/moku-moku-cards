package decks

import (
	"context"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (deck *Deck) Get() *errors.RestErr {
	err := mongo_db.DB.Collection("decks").FindOne(context.TODO(), bson.D{{"ID", deck.ID}}).Decode(&deck)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NotFoundError("deck not found")
		}
		log.Fatal(err)
	}
	return nil
}

func (deck *Deck) Delete() *errors.RestErr {
	result, err := mongo_db.DB.Collection("decks").DeleteOne(context.TODO(), bson.D{{"ID", deck.ID}})
	if result.DeletedCount == 0 {
		return errors.NotFoundError("deck not found")
	}
	if err != nil {
		panic(err)
	}
	return nil
}
