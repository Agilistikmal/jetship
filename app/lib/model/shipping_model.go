package model

type Shipping struct {
	ID          string    `json:"id" validate:"required"`
	Packages    []Package `json:"packages" validate:"required,min=1"`
	Statuses    []Status  `json:"statuses"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   uint64    `json:"created_at" validate:"required"`
	UpdatedAt   uint64    `json:"updated_at" validate:"required"`
}

type Status struct {
	Text      string `json:"text" validate:"required"`
	CreatedAt uint64 `json:"created_at" validate:"required"`
}
