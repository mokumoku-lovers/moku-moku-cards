package decks

import "go.mongodb.org/mongo-driver/bson/primitive"

type Deck struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name"`
	Cards   []string           `json:"cards"`
	Creator int64              `json:"creator"`
	Date    string             `json:"date"`
}
