package cards

import (
	"moku-moku-cards/utils/errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Front string             `json:"front" form:"front"`
	Back  string             `json:"back" form:"back"`
	Image string             `json:"image"`
}

// ValidateCard checks if all the data is correct
func (card *Card) ValidateCard() *errors.RestErr {
	// Card empty front side
	if card.Front == "" {
		return errors.BadRequest("missing card front side data")
	}
	// Card empty back side
	if card.Back == "" {
		return errors.BadRequest("missing card back side data")
	}

	return nil
}
