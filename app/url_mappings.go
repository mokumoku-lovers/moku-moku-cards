package app

import (
	"moku-moku-cards/controllers/cards"
	"moku-moku-cards/controllers/decks"
)

func mapUrls() {
	router.GET("/decks", decks.GetDecks)
	router.GET("/decks/:userID", decks.GetUserDecks)
	router.GET("/deck/:deckID", decks.GetDeck)
	router.POST("/deck", decks.CreateDeck)
	router.DELETE("/deck/:deckID", decks.DeleteDeck)
	router.PUT("/deck/:deckID", decks.UpdateDeck)

	router.GET("/card/:cardID", cards.GetCard)
	router.POST("/card", cards.PostCard)
	router.DELETE("/card/:cardID", cards.DeleteCard)
	router.PATCH("/card/:cardID", nil)
}
