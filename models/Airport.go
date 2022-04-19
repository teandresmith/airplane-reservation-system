package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Airport struct {
	ID				primitive.ObjectID		`bson:"_id" json:"_id"`
	Name      		*string   				`bson:"name" json:"name" validate:"required"`
	Acronym   		*string   				`bson:"acronym" json:"acronym" validate:"required"`
	Geopoint		Point					`bson:"geopoint" json:"geopoint" validate:"required"`
	City	  		*string					`bson:"city" json:"city" validate:"required"`
	Country	  		*string					`bson:"country" json:"country" validate:"required"`
	CreatedAt		time.Time				`bson:"createdAt" json:"createdAt"`
	UpdatedAt		time.Time				`bson:"updatedAt" json:"updatedAt"`
}

type Point struct{
	Latitude  		*float64  				`bson:"latitude" json:"latitude" validate:"required"`
	Longitude 		*float64				`bson:"longitude" json:"longitude" validate:"required"`
}