package model

type OrderItem struct {
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
}
