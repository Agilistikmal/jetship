package model

type Shipping struct {
	Packages    []Package `json:"packages"`
	Statuses    []Status  `json:"statuses"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   uint64    `json:"created_at"`
	UpdatedAt   uint64    `json:"updated_at"`
}

type Status struct {
	Text      string `json:"text"`
	CreatedAt uint64 `json:"created_at"`
}
