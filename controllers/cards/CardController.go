package cards

import (
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCard(c *gin.Context) {
	cardID, cardErr := strconv.ParseInt(c.Param("cardID"), 10, 64)
	if cardErr != nil {
		err := errors.BadRequest("card id should be a number")
		c.JSON(err.Status, err)
		return
	}
	card, getErr := services.GetCard(cardID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, card)
}

// PostCard controller creates a new card based on the data provided by the
// request body.
func PostCard(c *gin.Context) {
	var newCard cards.Card

	// Get card info from request body
	if err := c.ShouldBindJSON(&newCard); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	// Create card
	err := services.PostCard(newCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, newCard)
}
