package cards

import (
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
