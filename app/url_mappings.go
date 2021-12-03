package app

func mapUrls() {
	router.GET("/decks", nil)
	router.GET("/deck/:deckID", nil)
	router.POST("/deck", nil)
	router.DELETE("/deck/:deckID", nil)
	router.PATCH("/deck/:deckID", nil)

	router.GET("/card/:card_id", nil)
	router.POST("/card", nil)
	router.DELETE("/card/:cardID", nil)
	router.PATCH("/card/:cardID", nil)
}
