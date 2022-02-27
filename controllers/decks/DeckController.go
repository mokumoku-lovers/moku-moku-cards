package decks

import (
	"github.com/gin-gonic/gin"
	"moku-moku-cards/domain/decks"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
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
