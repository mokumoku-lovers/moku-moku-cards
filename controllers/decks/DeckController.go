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
