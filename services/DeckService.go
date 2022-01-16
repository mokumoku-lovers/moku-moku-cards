package services

import (
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/utils/errors"
)

func GetDeck(deckID int64) (*decks.Deck, *errors.RestErr) {
	result := &decks.Deck{ID: deckID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}