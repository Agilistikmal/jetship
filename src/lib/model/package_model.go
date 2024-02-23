package model

type Package struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Weight      int    `json:"weight"`
	Dimension   struct {
		Height int `json:"height"`
		Width  int `json:"width"`
		Length int `json:"length"`
	} `json:"dimension"`
	Sender    *Person `json:"sender"`
	Recipient *Person `json:"recipient"`
}
