package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Income struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Income int                `bson:"income" json:"income,omitempty"`
}
