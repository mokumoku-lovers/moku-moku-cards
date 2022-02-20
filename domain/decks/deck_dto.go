package decks

type Deck struct {
	ID      string  `json:"id" bson:"_id"`
	Name    string  `json:"name"`
	Cards   []int64 `json:"cards"`
	Creator int64   `json:"creator"`
	Date    string  `json:"date"`
}
