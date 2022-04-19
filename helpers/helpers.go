package helpers

import (
	"math"
	"strconv"

	"github.com/teandresmith/airplane-reservation-system/models"
)

func CalculateDistanceInMiles(pointA models.Point, pointB models.Point) float64 {

	pointALatitude := *pointA.Latitude * math.Pi / 180
	pointALongitude := *pointA.Longitude * math.Pi / 180

	pointBLatitude := *pointB.Latitude * math.Pi / 180
	pointBLongitude := *pointB.Longitude * math.Pi / 180

	// Apply Haversine Formula to calculate distance ( This will be used to calculate the price of flights )

	dlon := pointBLongitude - pointALongitude
	dlat := pointBLatitude - pointALatitude

	a := math.Pow(math.Sin(dlat / 2), 2) + math.Cos(pointALatitude) * math.Cos(pointBLatitude) * math.Pow(math.Sin(dlon / 2), 2)

	c := 2 * math.Asin(math.Sqrt(a))

	radius := 3956

	return(c * float64(radius))
}


/* 
	This function calculates the cost of a flight. This is not a realistic model
	and should not be taken seriously. It is only a formula to calculate a price of a ticket
	without having to manually determine the cost of a flight as there are too many options
	to handle.
*/
func CalculateFlightCost(distance float64, flightFareClass string, isInternational bool) float64 {

	// Using the cost of fuel per unit of measure in USD as of January 2022.
	// Cost of General Jet Fuel as of January 2022 is $2.14 p/gallon
	// Average Miles per gallon = 64 Miles
	// This is the additional cost on top of the basic prices of Flight fare
	/* 
		Price Formula =  ( Distance / Average Miles Per Gallon ) + Basic Fare Cost
		Domestic Fare Cost
		Economy = $25
		Business = $100
		First Class = $500
		
		International Fare Cost
		Economy = $300
		Business = $550
		First Class = $1500
	*/

	economy := 25
	business := 100
	firstClass := 500

	if isInternational {
		economy = 300
		business = 550
		firstClass = 1500
	}

	averageMilesPerGallon := 64
	var standardFareFee int

	switch flightFareClass {
	case "Economy":
		standardFareFee = economy
		break
	case "Business":
		standardFareFee = business
		break
	case "First Class":
		standardFareFee = firstClass
		break
	}

	cost := (distance / float64(averageMilesPerGallon)) * 2.14 + float64(standardFareFee)

	return cost
}

func CalculateDistanceAndCost(pointA models.Point, pointB models.Point, flightFareClass string, isInternational bool) interface{} {
	distance := CalculateDistanceInMiles(pointA, pointB)
	cost := CalculateFlightCost(distance, flightFareClass, isInternational)

	type Calculation struct{
		Distance float64 `bson:"distance" json:"distance"`
		Cost float64 `bson:"cost" json:"cost"`
	}

	return Calculation{
		Distance: distance,
		Cost: cost,
	}
}

func ToFloat(value string) (float64, error) {
	conversion, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return conversion, nil
}

func Round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(Round(num * output)) / output
}