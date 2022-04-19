package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc{
	return func(c *gin.Context) {

		// token := c.GetHeader("token")

		// if token == "" {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Error{
		// 		Message: "No token header provided",
		// 	})
		// 	return
		// }

		// claims, err := helpers.ValidateToken()
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, controllers.Error{
		// 		Message: "Token is invalid.",
		// 		Error: err.Error(),
		// 	})
		// 	return
		// }

		// c.Set("firstName", claims.FirstName)
		// c.Set("lastName", claims.LastName)
		// c.Set("flightNumber", claims.FlightNumber)
		// c.Set("reservationId", claims.ReservationID)
		// c.Set("date", claims.Date)


		c.Set("Permissions", "Granted")
	}
}