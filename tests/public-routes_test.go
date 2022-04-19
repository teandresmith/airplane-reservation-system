package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/teandresmith/airplane-reservation-system/controllers"
	"github.com/teandresmith/airplane-reservation-system/helpers"
	"github.com/teandresmith/airplane-reservation-system/models"
)

type Points struct{
	PointA models.Point	`bson:"pointA" json:"pointA"`
	PointB models.Point	`bson:"pointB" json:"pointB"`
}


func TestPublicPoints(t *testing.T) {

	t.Run("Returns Airports", func(t *testing.T) {
		response, err := http.Get("http://localhost:8080/api/airports")
		if err != nil {
			t.Errorf("Test Failed: %s", err.Error())
		}

		defer response.Body.Close()

		if response.Status != "200 OK" {
			t.Errorf("Got %s, wanting 200 OK", response.Status)
		}
	})

	t.Run("Return Distance and Cost", func(t *testing.T) {
		pointALat := -9.428
		pointALong := 160.054993

		pointBLat := -2.06189
		pointBLong := 147.423996

		points := Points{
			PointA: models.Point{
				Latitude: &pointALat,
				Longitude:  &pointALong,
			},
			PointB: models.Point{
				Latitude: &pointBLat,
				Longitude: &pointBLong,
			},
		}

		json_data, jsonErr := json.Marshal(points)
		if jsonErr != nil {
			t.Error(jsonErr.Error())
		}

		response, err := http.Post("http://localhost:8080/api/airport/distance-cost?fightFareClass=Economy&isInternational=true", "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			t.Errorf("Test Failed: %s", err.Error())
		}

		if response.Status != "200 OK" {
			t.Errorf("Got %s, wanting 200 OK", response.Status)
		}
	})

	t.Run("Return Correct Distance between Airports", func(t *testing.T) {
		pointALat := -9.428
		pointALong := 160.054993

		pointBLat := -2.06189
		pointBLong := 147.423996

		points := Points{
			PointA: models.Point{
				Latitude: &pointALat,
				Longitude:  &pointALong,
			},
			PointB: models.Point{
				Latitude: &pointBLat,
				Longitude: &pointBLong,
			},
		}

		json_data, jsonErr := json.Marshal(points)
		if jsonErr != nil {
			t.Error(jsonErr.Error())
		}

		response, err := http.Post("http://localhost:8080/api/airport/distance", "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			t.Errorf("Test Failed: %s", err.Error())
		}

		if response.Status != "200 OK" {
			t.Errorf("Got %s, wanted 200 OK", response.Status)
		}

		var res controllers.Response

		json.NewDecoder(response.Body).Decode(&res)

		checkDistance := helpers.CalculateDistanceInMiles(points.PointA, points.PointB)
		if res.Result != checkDistance {
			t.Errorf("Response Distance was %f, but wanted %f", res.Result, checkDistance)
		}

	})
	
	t.Run("Return Correct Flight Cost", func(t *testing.T) {
		response, err := http.Get("http://localhost:8080/api/airport/flight-cost?distance=1005.247704&flightFareClass=First+Class&isInternational=false")
		if err != nil {
			t.Errorf("Test Failed: %s", err.Error())
		}

		if response.Status != "200 OK" {
			t.Errorf("Wanted Status 200, but got %s", response.Status)
		}

		var res controllers.Response

		json.NewDecoder(response.Body).Decode(&res)

		
		checkedCost := helpers.CalculateFlightCost(1005.247704, "First Class", false)
		if res.Result != checkedCost {
			t.Errorf("Response Cost was %f, but got %f", res.Result, checkedCost)
		}
	})
}