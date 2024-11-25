package model

type Order struct {
	ID          string       `json:"id"`
	Items       []*OrderItem `json:"items"`
	TotalAmount float64      `json:"totalAmount"`
	Status      string       `json:"status"`
	CreatedAt   string       `json:"createdAt"`
}
