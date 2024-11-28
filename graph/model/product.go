package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Category *Category `json:"category"`
	Stock    int       `json:"stock"`
}

type MongoProduct struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Price      float64            `bson:"price"`
	CategoryID primitive.ObjectID `bson:"category_id"`
	Stock      int                `bson:"stock"`
}
