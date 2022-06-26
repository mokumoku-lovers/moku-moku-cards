package decks

import (
	"github.com/gin-gonic/gin"
	"github.com/mokumoku-lovers/moku-moku-oauth-go/oauth"
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"strconv"
)

func GetDeck(c *gin.Context) {
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
	deckID := c.Param("deckID")
	deck, getErr := services.GetDeck(deckID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, deck)
}

func DeleteDeck(c *gin.Context) {
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
	deckID := c.Param("deckID")

	_, deleteErr := services.DeleteDeck(deckID)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, "deck deleted")
}

// DeleteCardsFromDeck gets the specified deck data back
// looks for the card to be deleted from the deck and removes it
func DeleteCardsFromDeck(c *gin.Context) {
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
	deckID := c.Param("deckID")

	// Get the deck back with the list of current cards
	deck, getErr := services.GetDeck(deckID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	// Receive a slice of cardIDs to delete
	var cardsIDs []string
	if err := c.ShouldBindJSON(&cardsIDs); err != nil {
		bodyErr := errors.BadRequest("invalid json body")
		c.JSON(bodyErr.Status, bodyErr)
		return
	}

	// Modify the cards array from the deck by deleting the specified card
	services.DeleteCardsFromDeck(deck, cardsIDs)

	// Update the deck with the updated list of cards
	res, updateErr := services.PartialUpdateDeck(deckID, *deck)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateDeck(c *gin.Context) {
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
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
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
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
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
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
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
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
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
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

func UpdateDeckCards(c *gin.Context) {
	requestErr := oauth.AuthenticateRequest(c.Request)
	if requestErr != nil {
		c.JSON(requestErr.Status, requestErr)
		return
	}
	deckID := c.Param("deckID")
	var deck decks.Deck
	if err := c.ShouldBindJSON(&deck); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, updateErr := services.UpdateDeckCards(deckID, deck)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
