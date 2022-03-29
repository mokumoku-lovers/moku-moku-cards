package decks

import (
	"github.com/gin-gonic/gin"
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"strconv"
)

func GetDeck(c *gin.Context) {
	deckID := c.Param("deckID")
	deck, getErr := services.GetDeck(deckID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, deck)
}

func DeleteDeck(c *gin.Context) {
	deckID := c.Param("deckID")

	_, deleteErr := services.DeleteDeck(deckID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, "deck deleted")
}

// DeleteCardFromDeck gets the specified deck data back
// looks for the card to be deleted from the deck and removes it
func DeleteCardFromDeck(c *gin.Context) {
	deckID := c.Param("deckID")
	cardID, _ := strconv.ParseInt(c.Param("cardID"), 10, 64)

	// Get the deck back with the list of current cards
	deck, getErr := services.GetDeck(deckID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	// Modify the cards array from the deck by deleting the specified card
	services.DeleteCardFromDeck(deck, cardID)

	// Update the deck with the updated list of cards
	res, updateErr := services.PartialUpdateDeck(deckID, *deck)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateDeck(c *gin.Context) {
	var deck decks.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateDeck(deck)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateDeck(c *gin.Context) {
	deckID := c.Param("deckID")
	var deck decks.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, updateErr := services.UpdateDeck(deckID, deck)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func PartialUpdateDeck(c *gin.Context) {
	deckID := c.Param("deckID")
	var deck decks.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, updateErr := services.PartialUpdateDeck(deckID, deck)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//get all decks in db
func GetDecks(c *gin.Context) {
	decks, getErr := services.GetDecks()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, decks)
}

// GetUserDecks retrieves all the decks from the
// specified user
func GetUserDecks(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		badRequest := errors.BadRequest("invalid userID")
		c.JSON(badRequest.Status, badRequest)
		return
	}

	decks, getErr := services.GetUserDecks(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, decks)
}
