package entity

type Value struct {
	Name     string  `json:"name"`
	Amount   float32 `json:"amount"`
	Invoiced bool    `json:"invoiced"`
}
