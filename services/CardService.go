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
