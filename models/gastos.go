package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gasto struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Year        int                `bson:"year" json:"year,omitempty"`
	Month       string             `bson:"month" json:"month,omitempty"`
	ExpenseName string             `bson:"expenseName" json:"expenseName,omitempty"`
	Expense     int                `bson:"expense" json:"expense,omitempty"`
}
