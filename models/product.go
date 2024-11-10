package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product defines the structure of product data
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Brand       string             `json:"brand"`
	Category    string             `json:"category"`
	Price       float64            `json:"price"`
	Stock       int                `json:"stock"`
	ImagePath   string             `json:"imagePath"` // Path to the product image
}
