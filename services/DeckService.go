package services

import (
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/utils/errors"
)

func GetDeck(deckID string) (*decks.Deck, *errors.RestErr) {
	result := &decks.Deck{ID: deckID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func GetDecks() ([]decks.Deck, *errors.RestErr) {
	result, err := decks.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserDecks(userID int64) ([]decks.Deck, *errors.RestErr) {
	result, err := decks.GetAllUserDecks(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
