package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reserveration struct {
	ID        		primitive.ObjectID	`bson:"_id" json:"_id"`
	FirstName 		*string  			`bson:"firstName" json:"firstName" validate:"required"`
	LastName 		*string  			`bson:"lastName" json:"lastName" validate:"required"`
	Email     		*string  			`bson:"email" json:"email" validate:"required"`
	TotalPrice		*float64			`bson:"totalPrice" json:"totalPrice" validate:"required"`
	Flights   		[]Flight 			`bson:"flight" json:"flight" validate:"required"`
	IsPaid    		bool     			`bson:"isPaid" json:"isPaid"`
	CreatedAt		time.Time			`bson:"createdAt" json:"createdAt"`
	UpdatedAt		time.Time			`bson:"updatedAt" json:"updatedAt"`
}
