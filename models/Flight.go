package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID						primitive.ObjectID	`bson:"_id" json:"_id"`
	Airplane				Airplane			`bson:"airplane" json:"airplane"`
	Seat					SelectedSeat		`bson:"seat" json:"seat"`		
	Departure				Airport				`bson:"departure" json:"departure" validate:"required"`
	Destination				Airport				`bson:"destination" json:"destination" validate:"required" `
	FlightDuration 			time.Time 			`bson:"flightDuration" json:"flightDuration" validate:"required"`
	DepartureTime			*string				`bson:"departureTime" json:"departureTime" validate:"required"`
	ArrivalTime				*string				`bson:"arrivalTime" json:"arrivalTime" validate:"required"`
	Price					*float64			`bson:"price" json:"price" validate:"required"`
	CreatedAt				time.Time			`bson:"createdAt" json:"createdAt"`
	UpdatedAt				time.Time			`bson:"updatedAt" json:"updatedAt"`	
}

type SelectedSeat struct {
	SeatNumber				*string				`bson:"seatNumber" json:"seatNumber"`	
	SeatType				*string				`bson:"seatType" json:"seatType"`	
}








