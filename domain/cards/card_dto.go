package cards

import "moku-moku-cards/utils/errors"

type Card struct {
	ID    int64  `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	Image string `json:"image"`
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
