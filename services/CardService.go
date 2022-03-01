package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/utils/errors"
)

func GetCard(cardID string) (*cards.Card, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return nil, errors.BadRequest("invalid card ID")
	}
	result := &cards.Card{ID: objectID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// PostCard validates the received card information
// and saves it in the DB
func PostCard(card *cards.Card) *errors.RestErr {
	// Validate card information
	if err := card.ValidateCard(); err != nil {
		return err
	}

	card.ID = primitive.NewObjectID()
	// Save card to the DB
	if err := card.Post(); err != nil {
		return err
	}

	return nil
}

func DeleteCard(cardID string) (string, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(cardID)
	if err != nil {
		return "", errors.BadRequest("invalid card ID")
	}
	result := &cards.Card{ID: objectID}
	if err := result.Delete(); err != nil {
		return "", err
	}
	return "card deleted", nil
}
