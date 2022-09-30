package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/utils/errors"
	"moku-moku-cards/utils/slices"
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
	objectID, parseIDError := primitive.ObjectIDFromHex(cardID)
	if parseIDError != nil {
		return "", errors.BadRequest("invalid card ID")
	}
	result := &cards.Card{ID: objectID}
	if deleteCardError := result.Delete(); deleteCardError != nil {
		return "", deleteCardError
	}
	allDecks, getDecksError := GetDecks()
	if getDecksError != nil {
		return "", getDecksError
	}
	for _, deck := range allDecks {
		for cardIndex := range deck.Cards {
			if deck.Cards[cardIndex] == cardID {
				deck.Cards = slices.RemoveIndex(deck.Cards, cardIndex)
				deck.UpdateField("Cards")
				break
			}
		}
	}
	return "card deleted", nil
}

func PartialUpdateCard(cardID string, card cards.Card) (int64, *errors.RestErr) {
	objectId, idErr := primitive.ObjectIDFromHex(cardID)
	if idErr != nil {
		return 0, errors.BadRequest("Invalid ID")
	}
	card.ID = objectId
	result, updateErr := card.PartialUpdate()
	if updateErr != nil {
		return 0, updateErr
	}
	return result, nil
}
