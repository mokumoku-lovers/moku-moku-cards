package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"moku-moku-cards/controllers/cards"
	"moku-moku-cards/controllers/decks"
	"net/http"
)

func mapUrls() {
	router.GET("/decks", decks.GetDecks)
	router.GET("/decks/:userID", decks.GetUserDecks)
	router.GET("/deck/:deckID", decks.GetDeck)
	router.POST("/deck", decks.CreateDeck)
	router.DELETE("/deck/:deckID", decks.DeleteDeck)
	router.DELETE("/deck/:deckID/cards", decks.DeleteCardsFromDeck)
	router.PUT("/deck/:deckID", decks.UpdateDeck)
	router.PATCH("/deck/:deckID", decks.PartialUpdateDeck)
	router.PATCH("/deck/:deckID/cards", decks.UpdateDeckCards)

	router.GET("/card/:cardID", cards.GetCard)
	router.POST("/card", cards.PostCard)
	router.DELETE("/card/:cardID", cards.DeleteCard)
	router.PATCH("/card/:cardID", cards.PartialUpdateCard)

	// Swagger documentation
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml", Title: "Moku-Moku-Cards"}
	swg := middleware.Redoc(opts, nil)
	router.GET("/docs", gin.WrapH(swg))
	router.GET("/swagger.yml", gin.WrapH(http.FileServer(http.Dir("./"))))
}
