package decks

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/docs"
	"moku-moku-cards/utils/errors"
	"moku-moku-cards/utils/slices"
	"reflect"
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
	result, err := mongo_db.DB.Collection("decks").DeleteOne(context.TODO(), bson.D{{"_id", deck.ID}})
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

func (deck *Deck) UpdateField(fieldName string) (int64, *errors.RestErr) {
	metaValue := reflect.ValueOf(deck).Elem()
	field := metaValue.FieldByName(fieldName)
	if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
		bName, bErr := docs.FieldToBson(fieldName)
		if bErr != nil {
			return 0, errors.BadRequest(bErr.Error())
		}
		_, err := mongo_db.DB.Collection("decks").UpdateOne(context.TODO(), bson.M{"_id": deck.ID},
			bson.D{
				{"$set", bson.M{bName: field.Interface()}},
			})
		if err != nil {
			return 0, errors.InternalServerError("failed updating document")
		}
	}
	return 1, nil
}

func (deck *Deck) PartialUpdate() (int64, *errors.RestErr) {
	for _, fieldName := range []string{"Name", "Cards", "Creator", "Date"} {
		_, err := deck.UpdateField(fieldName)
		if err != nil {
			return 0, errors.InternalServerError("failed updating document field " + fieldName)
		}
	}
	return 1, nil
}

func (deck *Deck) UpdateCards() (int64, *errors.RestErr) {
	var original Deck
	findErr := mongo_db.DB.Collection("decks").FindOne(context.TODO(), bson.D{{"_id", deck.ID}}).Decode(&original)
	if findErr != nil {
		if findErr == mongo.ErrNoDocuments {
			return 0, errors.NotFoundError("deck not found")
		}
		log.Fatal(findErr)
	}
	deck.Cards = slices.Deduplicate(append(original.Cards, deck.Cards...))
	_, updateErr := deck.UpdateField("Cards")
	if updateErr != nil {
		return 0, errors.InternalServerError("failed updating document field Cards")
	}
	return 1, nil
}
