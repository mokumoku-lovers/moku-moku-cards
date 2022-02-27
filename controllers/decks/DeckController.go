package decks

import (
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
