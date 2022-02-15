package decks

import (
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDeck(c *gin.Context) {
	deckID, deckErr := strconv.ParseInt(c.Param("deckID"), 10, 64)
	if deckErr != nil {
		err := errors.BadRequest("deck id should be a number")
		c.JSON(err.Status, err)
		return
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
