package cards

type Card struct {
	ID    int64  `json:"id"`
	Front string `json:"front"`
	Back  string `json:"back"`
	Image string `json:"image"`
}
