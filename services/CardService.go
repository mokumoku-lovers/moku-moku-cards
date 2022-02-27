package services

import (
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/utils/errors"
)

func GetCard(cardID int64) (*cards.Card, *errors.RestErr) {
	result := &cards.Card{ID: cardID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// PostCard validates the received card information
// and saves it in the DB
func PostCard(card cards.Card) *errors.RestErr {
	// Validate card information
	if err := card.ValidateCard(); err != nil {
		return err
	}

	// Save card to the DB
	if err := card.Post(); err != nil {
		return err
	}

	return nil
}
