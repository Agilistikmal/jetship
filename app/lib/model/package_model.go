package model

type Package struct {
	ID        string  `json:"id" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Details   string  `json:"details" validate:"required,min=8"`
	Weight    float32 `json:"weight" validate:"required,min=0.1"`
	Dimension struct {
		Height float32 `json:"height"`
		Width  float32 `json:"width"`
		Length float32 `json:"length"`
	} `json:"dimension"`
	Sender    *Person `json:"sender" validate:"required"`
	Recipient *Person `json:"recipient" validate:"required"`
}
