package public

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teandresmith/airplane-reservation-system/controllers"
	"github.com/teandresmith/airplane-reservation-system/database"
	"github.com/teandresmith/airplane-reservation-system/helpers"
	"github.com/teandresmith/airplane-reservation-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var airportCollection *mongo.Collection = database.OpenCollection(database.Client, "Airports")

type Points struct{
	PointA models.Point	`bson:"pointA" json:"pointA"`
	PointB models.Point	`bson:"pointB" json:"pointB"`
}

func GetDistanceBetweenAirport() gin.HandlerFunc{
	return func(c * gin.Context) {

		var data Points

		if bindErr := c.BindJSON(&data); bindErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while binding request body data",
				Error: bindErr.Error(),
			})
			return
		}

		distance := helpers.CalculateDistanceInMiles(data.PointA, data.PointB)

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Successful",
			Result: distance,
		})
	}
}

func GetFlightCost() gin.HandlerFunc{
	return func(c *gin.Context) {

		distance, distanceErr := strconv.ParseFloat(c.Query("distance"), 64)
		if distanceErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while converting query params.",
				Error: distanceErr.Error(),
			})
		}

		flightFareClass := c.Query("flightFareClass")
		if flightFareClass == "" {
			flightFareClass = "Economy"
		}

		isInternational, isInternationalErr := strconv.ParseBool(c.Query("isInternational"))
		if isInternationalErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while converting query params.",
				Error: isInternationalErr.Error(),
			})
		}

		cost := helpers.CalculateFlightCost(distance, flightFareClass, isInternational)

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Success",
			Result: cost,
		})
	}
}

func GetDistanceAndCost() gin.HandlerFunc{
	return func(c *gin.Context) {
		flightFareClass := c.Query("flightFareClass")
		if flightFareClass == "" {
			flightFareClass = "Economy"
		}

		isInternational, isInternationalErr := strconv.ParseBool(c.Query("isInternational"))
		if isInternationalErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while converting query params.",
				Error: isInternationalErr.Error(),
			})
		}
		
		var data Points

		if bindErr := c.BindJSON(&data); bindErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while binding request body data",
				Error: bindErr.Error(),
			})
			return
		}


		info := helpers.CalculateDistanceAndCost(data.PointA, data.PointB, flightFareClass, isInternational)

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Success",
			Result: info,
		})
	}
}

func GetAirports() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var aiports []models.Airport

		query, queryErr := airportCollection.Find(ctx, bson.M{})
		defer cancel()
		if queryErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while querying the Airports Collection",
				Error: queryErr.Error(),
			})
			return
		}

		iterateErr := query.All(ctx, &aiports)
		defer cancel() 
		if iterateErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while iterating through the query results",
				Error: iterateErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Query Successful",
			Result: aiports,
		})
	}
}