package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teandresmith/airplane-reservation-system/controllers/private"
)

func PrivateRoutes(incomingRoutes *gin.Engine) {
	/* 
		All routes within this file will be protected by a JWT. There will be no login process, but rather
		a user must provide the correct information to access a reservation.

		These fields will be:
			Reservation Date
			Flight Number
			Reservation Number
			First Name
			Last Name
	*/
	incomingRoutes.GET("/api/reservation/:reservationid", private.GetReservation())
	incomingRoutes.POST("/api/reservation/:reservationid", private.CreateReservation())
	incomingRoutes.PATCH("/api/reservation/:reservationid", private.UpdateReservation())
	incomingRoutes.DELETE("/api/reservation/:reservationid", private.DeleteReservation())
}