package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Level     string             `json:"level" bson:"level"`
	Message   string             `json:"message" bson:"message"`
	CreatedOn primitive.DateTime `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
}
