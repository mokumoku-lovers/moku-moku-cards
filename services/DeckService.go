package services

import (
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/utils/errors"
)

func GetDeck(deckID string) (*decks.Deck, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(deckID)
	if err != nil {
		return nil, errors.NotFoundError("Invalid ID")
	}
	result := &decks.Deck{ID: objectID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateDeck(deck decks.Deck) (string, *errors.RestErr) {
	// May need validation here.
	res, err := deck.Save()
	if err != nil {
		return "", err
	}
	return res, nil
}
