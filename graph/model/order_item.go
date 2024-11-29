package model

type OrderItem struct {
	Product  *Product `json:"product"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
}

type MongoOrderItem struct {
	ProductID string  `bson:"productID"`
	Quantity  int     `bson:"quantity"`
	Price     float64 `bson:"price"`
}
