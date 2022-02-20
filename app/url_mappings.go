package app

import (
	"moku-moku-cards/controllers/cards"
	"moku-moku-cards/controllers/decks"
)

func mapUrls() {
	router.GET("/decks", decks.GetDecks)
	router.GET("/deck/:deckID", decks.GetDeck)
	router.POST("/deck", nil)
	router.DELETE("/deck/:deckID", nil)
	router.PATCH("/deck/:deckID", nil)

	router.GET("/card/:cardID", cards.GetCard)
	router.POST("/card", cards.PostCard)
	router.DELETE("/card/:cardID", nil)
	router.PATCH("/card/:cardID", nil)
}
