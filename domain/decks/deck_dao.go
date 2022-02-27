package decks

import (
	"context"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/docs"
	"moku-moku-cards/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (deck *Deck) Get() *errors.RestErr {
	err := mongo_db.DB.Collection("decks").FindOne(context.TODO(), bson.D{{"_id", deck.ID}}).Decode(&deck)
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
func GetAllUserDecks(userID int64) ([]Deck, *errors.RestErr) {
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

func (deck *Deck) Save() (primitive.ObjectID, *errors.RestErr) {
	res, err := mongo_db.DB.Collection("decks").InsertOne(context.TODO(), deck)
	if err != nil {
		return primitive.NilObjectID, errors.NotFoundError("failed writing document")
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (deck *Deck) Update() (int64, *errors.RestErr) {
	updateDoc, docErr := docs.InterfaceToDoc(deck)
	if docErr != nil {
		return 0, errors.BadRequest("invalid Doc")
	}
	result, err := mongo_db.DB.Collection("decks").UpdateOne(context.TODO(), bson.M{"_id": deck.ID},
		bson.D{
			{"$set", updateDoc},
		})
	if err != nil {
		return 0, errors.NotFoundError("failed updating document")
	}
	return result.ModifiedCount, nil
}
