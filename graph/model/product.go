package model

type Product struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Category *Category `json:"category"`
	Stock    int       `json:"stock"`
}
