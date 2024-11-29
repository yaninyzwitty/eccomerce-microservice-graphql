package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          string       `json:"id"`
	Items       []*OrderItem `json:"items"`
	TotalAmount float64      `json:"totalAmount"`
	Status      string       `json:"status"`
	CreatedAt   string       `json:"createdAt"`
}

// Order represents an order in the MongoDB collection.
type MongoOrder struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TotalAmount float64            `bson:"totalAmount"`
	Status      string             `bson:"status"`
	Items       []*MongoOrderItem  `bson:"items"`
	CreatedAt   string             `bson:"createdAt"`
}
