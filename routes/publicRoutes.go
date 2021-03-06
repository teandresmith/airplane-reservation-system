package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teandresmith/airplane-reservation-system/controllers/public"
	"github.com/teandresmith/airplane-reservation-system/stripe"
)

func PublicRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/airport/distance", public.GetDistanceBetweenAirport())
	incomingRoutes.GET("/api/airport/flight-cost", public.GetFlightCost())
	incomingRoutes.POST("/api/airport/distance-cost", public.GetDistanceAndCost())

	//
	incomingRoutes.GET("/api/airports", public.GetAirports())

	// Create Payment Intent
	incomingRoutes.POST("/api/create-payment-intent", stripe.CreatePaymentIntent())
}