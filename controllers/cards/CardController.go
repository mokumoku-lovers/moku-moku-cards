package cards

import (
	"crypto/sha256"
	"encoding/hex"
	"moku-moku-cards/domain/cards"
	"moku-moku-cards/services"
	"moku-moku-cards/utils/errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	//Get card info from request body
	if err := c.ShouldBind(&newCard); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	//get image file
	file, _ := c.FormFile("file")
	if file != nil {
		fileType := file.Header.Get("Content-Type")
		if fileType != "image/jpeg" {
			c.JSON(http.StatusBadRequest, errors.BadRequest("file must be of type image"))
		}
		name := file.Filename
		hashedName := sha256.Sum256([]byte(name))
		hashedNameString := hex.EncodeToString(hashedName[:][:])
		newCard.Image = hashedNameString

		//write file to basePath
		basePath := "/MokuMoku/card_images/"
		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			//create directory
			os.MkdirAll(basePath, 0700)
		}
		saveErr := c.SaveUploadedFile(file, basePath+hashedNameString+".png")
		if saveErr != nil {
			c.JSON(http.StatusInternalServerError, errors.InternalServerError("file could not be saved"))
		}

	}

	// Create card
	err := services.PostCard(&newCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
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

func PartialUpdateCard(c *gin.Context) {
	cardID := c.Param("cardID")
	var card cards.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		restErr := errors.BadRequest("invalid json body")
		c.JSON(restErr.Status, restErr)
	}
	result, updateErr := services.PartialUpdateCard(cardID, card)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
