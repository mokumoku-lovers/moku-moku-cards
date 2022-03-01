package cards

import (
	"github.com/gin-gonic/gin"
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
)

func GetCard(c *gin.Context) {
	cardID := c.Param("cardID")
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

	// TODO: Save the image and set newCard.Image

	// Create card
	err := services.PostCard(&newCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, newCard)
}

func DeleteCard(c *gin.Context) {
	cardID := c.Param("cardID")
	msg, getErr := services.DeleteCard(cardID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, msg)
}
