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

func GetAll() ([]Deck, *errors.RestErr) {
	result, err := mongo_db.DB.Collection("decks").Find(context.TODO(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFoundError("no decks found")
		}
		log.Fatal(err)
	}
	var res []Deck
	//TODO: error- cannot decode string into an int, fix deck id
	if err = result.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}
	return res, nil
}

// GetAllUserDecks retrieves a user's decks from the DB
func GetAllUserDecks(userID string) ([]Deck, *errors.RestErr) {
	result, err := mongo_db.DB.Collection("decks").Find(context.TODO(), bson.M{"creator": userID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NotFoundError("no decks found for the specified user")
		}
		log.Fatal(err)
	}
	var res []Deck
	//TODO: error- cannot decode string into an int, fix deck id
	if err = result.All(context.Background(), &res); err != nil {
		log.Fatal(err)
	}
	return res, nil
}
