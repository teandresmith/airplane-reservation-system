package private

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/teandresmith/airplane-reservation-system/controllers"
	"github.com/teandresmith/airplane-reservation-system/database"
	"github.com/teandresmith/airplane-reservation-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var reservationCollection *mongo.Collection = database.OpenCollection(database.Client, "Reservations")
var validate = validator.New()


func GetReservation() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		permissions, exists := c.Get("Permissions")
		if !exists {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		if permissions != "Granted" {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		reservationId := c.Param("reservationid")

		var reservation models.Reserveration

		findErr := reservationCollection.FindOne(ctx, bson.M{"_id": reservationId}).Decode(&reservation)
		defer cancel()
		if findErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while querying the Reservation collection",
				Error: findErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Query Successful",
			Result: reservation,
		})
	}
}

func CreateReservation() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var newReservation models.Reserveration
		
		if bindErr := c.BindJSON(&newReservation); bindErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while binding the request body data",
				Error: bindErr.Error(),
			})
			return
		}

		if validateErr := validate.Struct(&newReservation); validateErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while validating request body data into required schema.",
				Error: validateErr.Error(),
			})
			return
		}

		newReservation.ID = primitive.NewObjectID()
		newReservation.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newReservation.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newReservation.IsPaid = false

		insert, insertErr := reservationCollection.InsertOne(ctx, newReservation)
		defer cancel()
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while inserting a new document in the Reservation Collection",
				Error: insertErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Insertion Successful",
			Result: insert,
		})
	}
}

func UpdateReservation() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		permissions, exists := c.Get("Permissions")
		if !exists {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		if permissions != "Granted" {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		reservationId := c.Param("reservationid")

		var reservation models.Reserveration
		var newReservation bson.D
		
		if bindErr := c.BindJSON(&reservation); bindErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while binding the request body data",
				Error: bindErr.Error(),
			})
			return
		}

		if reservation.FirstName != nil {
			newReservation = append(newReservation, bson.E{Key: "firstName", Value: reservation.FirstName})
		}

		if reservation.LastName != nil {
			newReservation = append(newReservation, bson.E{Key: "lastName", Value: reservation.LastName})
		}

		if reservation.Email != nil {
			newReservation = append(newReservation, bson.E{Key: "email", Value: reservation.Email})
		}

		if reservation.TotalPrice != nil {
			newReservation = append(newReservation, bson.E{Key: "totalPrice", Value: reservation.TotalPrice})
		}

		if reservation.Flights != nil {
			newReservation = append(newReservation, bson.E{Key: "flights", Value: reservation.Flights})
		}

		if reservation.IsPaid {
			newReservation = append(newReservation, bson.E{Key: "isPaid", Value: reservation.IsPaid})
		}

		filter := bson.M{"_id": reservationId}
		opts := options.Update().SetUpsert(true)
		update := bson.D{{Key: "$set", Value: newReservation}}


		collUpdated, updateErr := reservationCollection.UpdateOne(ctx, filter, update, opts)
		defer cancel()
		if updateErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while updating a document in the Reservation Collection",
				Error: updateErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Update Successful",
			Result: collUpdated,
		})
		
	}
}

func DeleteReservation() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		permissions, exists := c.Get("Permissions")
		if !exists {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		if permissions != "Granted" {
			c.JSON(http.StatusUnauthorized, controllers.Error{
				Message: "Unauthorized User",
			})
			return
		}

		reservationId := c.Param("reservationid")

		delete, deleteErr := reservationCollection.DeleteOne(ctx, bson.M{"_id": reservationId})
		defer cancel()
		if deleteErr != nil {
			c.JSON(http.StatusInternalServerError, controllers.Error{
				Message: "There was an error while deleting a document in the Reservations Collection",
				Error: deleteErr.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Deletion Successful",
			Result: delete,
		})
	}
}