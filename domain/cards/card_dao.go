package cards

import (
	"context"
	"log"
	"moku-moku-cards/datasources/mongo_db"
	"moku-moku-cards/utils/docs"
	"moku-moku-cards/utils/errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func (card *Card) Get() *errors.RestErr {
	err := mongo_db.DB.Collection("cards").FindOne(context.TODO(), bson.D{{"_id", card.ID}}).Decode(&card)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NotFoundError("card not found")
		}
		log.Fatal(err)
	}
	return nil
}

// Post stores the newly create card to the DB
func (card *Card) Post() *errors.RestErr {
	_, err := mongo_db.DB.Collection("cards").InsertOne(
		context.Background(),
		card,
	)

	if err != nil {
		log.Fatal(err)
		return errors.InternalServerError(err.Error())
	}

	// TODO: Get inserted card generated ID

	return nil
}

// Delete the specified card from the DB
func (card *Card) Delete() *errors.RestErr {
	one, err := mongo_db.DB.Collection("cards").DeleteOne(
		context.Background(),
		bson.D{{"_id", card.ID}},
	)

	if err != nil {
		return errors.BadRequest("bad request")
	}

	if one.DeletedCount == 0 {
		return errors.NotFoundError("card not found")
	}

	return nil
}

func (card *Card) PartialUpdate() (int64, *errors.RestErr) {
	metaValue := reflect.ValueOf(card).Elem()
	for _, name := range []string{"Front", "Back", "Image"} {
		field := metaValue.FieldByName(name)
		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			bName, bErr := docs.FieldToBson(name)
			if bErr != nil {
				return 0, errors.BadRequest(bErr.Error())
			}
			_, err := mongo_db.DB.Collection("cards").UpdateOne(
				context.TODO(),
				bson.M{"_id": card.ID},
				bson.D{
					{"$set", bson.M{bName: field.Interface()}},
				})
			if err != nil {
				return 0, errors.InternalServerError("failed updating document")
			}
		}
	}
	return 1, nil
}
