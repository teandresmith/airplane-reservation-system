package stripe

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/teandresmith/airplane-reservation-system/controllers"
)

type Payment struct{
	Amount int64 `bson:"amount" json:"amount"`
	Currency string	`bson:"currency" json:"currency"`
}

func CreatePaymentIntent() gin.HandlerFunc{
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var payment Payment
		
		if bindErr := c.BindJSON(&payment); bindErr != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was an error while binding the request body data",
				Error: bindErr.Error(),
			})
		}

		params := &stripe.PaymentIntentParams{
			Amount: stripe.Int64(payment.Amount),
			Currency: stripe.String(payment.Currency),
		}

		pi, err := paymentintent.New(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, controllers.Error{
				Message: "There was error while creating a new payment intent",
				Error: err.Error(),
			})
		}

		c.JSON(http.StatusOK, controllers.Response{
			Message: "Successfully created new payment intent",
			Result: pi.ClientSecret,
		})
	}
}

