package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DeleteDeck(deckID string) (*decks.Deck, *errors.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(deckID)
	if err != nil {
		return nil, errors.NotFoundError("Invalid ID")
	}
	result := &decks.Deck{ID: objectID}
	if err := result.Delete(); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCardsFromDeck removes the specified card from the card list of the deck
func DeleteCardsFromDeck(deck *decks.Deck, cardsIDs []string) {
	// TODO: Could be improved using a hashmap
	// Loop the slice of cards to be deleted
	for _, toDelete := range cardsIDs {
		// Loop and find the card to delete
		for i, card := range deck.Cards {
			// Once found, get left slice and right slice
			if card == toDelete {
				left := deck.Cards[:i]
				right := deck.Cards[i+1:]
				deck.Cards = append(left, right...)
			}
		}
	}
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

func CreateDeck(deck decks.Deck) (primitive.ObjectID, *errors.RestErr) {
	// May need validation here.
	deck.ID = primitive.NewObjectID()
	res, err := deck.Save()
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res, nil
}

func UpdateDeck(deckID string, deck decks.Deck) (int64, *errors.RestErr) {
	objectID, idErr := primitive.ObjectIDFromHex(deckID)
	if idErr != nil {
		return 0, errors.BadRequest("Invalid ID")
	}
	deck.ID = objectID
	result, updateErr := deck.Update()
	if updateErr != nil {
		return 0, updateErr
	}
	return result, nil
}

func PartialUpdateDeck(deckID string, deck decks.Deck) (int64, *errors.RestErr) {
	objectID, idErr := primitive.ObjectIDFromHex(deckID)
	if idErr != nil {
		return 0, errors.BadRequest("Invalid ID")
	}
	deck.ID = objectID
	result, updateErr := deck.PartialUpdate()
	if updateErr != nil {
		return 0, updateErr
	}
	return result, nil
}

func UpdateDeckCards(deckID string, deck decks.Deck) (int64, *errors.RestErr) {
	objectID, idErr := primitive.ObjectIDFromHex(deckID)
	if idErr != nil {
		return 0, errors.BadRequest("Invalid ID")
	}
	deck.ID = objectID
	result, updateErr := deck.UpdateCards()
	if updateErr != nil {
		return 0, updateErr
	}
	return result, nil
}
