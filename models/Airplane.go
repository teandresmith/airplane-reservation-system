package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Airplane struct {
	ID			  	primitive.ObjectID	`bson:"_id" json:"_id"`
	MakerName     	*string 			`bson:"makerName" json:"makerName" validate:"required"`
	AirplaneModel 	*string 			`bson:"airplaneModel" json:"airplaneModel" validate:"required"`
	FlightNumber  	*string 			`bson:"flightNumber" json:"flightNumbmer" validate:"required"`
	SeatInfo      	Seats   			`bson:"seatInfo" json:"seatInfo" gorm:"embedded"`
	CreatedAt		time.Time			`bson:"createdAt" json:"createdAt"`
	UpdatedAt		time.Time			`bson:"updatedAt" json:"updatedAt"`
}

type Seats struct {
	FirstClass uint
	Business   uint
	Economy    uint
	Total      uint
}