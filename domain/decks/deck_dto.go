package decks

type Deck struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	Cards   []int64 `json:"cards"`
	Creator int64   `json:"creator"`
	Date    string  `json:"date"`
}
