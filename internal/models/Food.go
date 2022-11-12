package models

import (
"go.mongodb.org/mongo-driver/bson/primitive"
"time"
)

type Food struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OrderDate  time.Time          `json:"order_date,omitempty" bson:"order_date,omitempty"`
	Region     string             `json:"region" bson:"region,omitempty"`
	City       string             `json:"city" bson:"city,omitempty"`
	Category   string             `json:"category" bson:"category,omitempty"`
	Product    string             `json:"product" bson:"product,omitempty"`
	Quantity   int64              `json:"quantity" bson:"quantity,omitempty"`
	UnitPrice  float64            `json:"unit_price" bson:"unit_price,omitempty"`
	TotalPrice float64            `json:"total_price" bson:"total_price,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
