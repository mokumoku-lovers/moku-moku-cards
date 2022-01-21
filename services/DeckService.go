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

func CreateDeck(deck decks.Deck) (string, *errors.RestErr) {
	// May need validation here.
	res, err := deck.Save()
	if err != nil {
		return "", err
	}
	return res, nil
}
