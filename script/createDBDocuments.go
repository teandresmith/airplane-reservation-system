package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/teandresmith/airplane-reservation-system/database"
	"github.com/teandresmith/airplane-reservation-system/helpers"
	"github.com/teandresmith/airplane-reservation-system/models"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func main() {
    LoadExcelData()
}

func LoadExcelData() {
    fmt.Println("Inserting Excel Data into Collection...")
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    
    deleteDataErr := database.OpenCollection(database.Client, "Airports").Drop(ctx)
    defer cancel()
    if deleteDataErr != nil {
        log.Panic(deleteDataErr)
    }
    airportCollection := database.OpenCollection(database.Client, "Airports")
    
	f, err := excelize.OpenFile("./airports.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
		
	var airports []models.Airport

    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("LargeAirports")
    if err != nil {
        fmt.Println(err)
        return
    }
    for index, row := range rows {
		if index == 0 {
			continue
		}
		name, country, city, acronym := row[1], row[4], row[5], row[6]
        
        latitude, err := helpers.ToFloat(row[2])
        if err != nil {
            return
        }

        longitude, err := helpers.ToFloat(row[3])
        if err != nil {
            return
        }


        createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
        updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		airport := models.Airport{
			ID:        primitive.NewObjectID(),
			Name:      &name,
			Acronym:   &acronym,
			Geopoint:  models.Point{Latitude: &latitude, Longitude: &longitude},
			City:      &city,
			Country:   &country,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}
        airports = append(airports, airport)
		
    }

    

    airportInterface := structToInterfaceArray(airports)

    _, insertErr := airportCollection.InsertMany(ctx, airportInterface)
    defer cancel()
    if insertErr != nil {
        log.Panic(err)
    }

    fmt.Println("Insertion Complete")

}

func structToInterfaceArray(structArray []models.Airport) []interface{} {
    var interfaceArray []interface{}
    
    for _, item := range structArray {
        airportToInterface := bson.D{
            {Key:"_id", Value: item.ID},{Key: "name", Value: item.Name},{Key: "acronym", Value: item.Acronym},{Key: "geopoint", Value: item.Geopoint},
            {Key: "city", Value: item.City}, {Key: "country", Value: item.Country}, {Key: "createdAt", Value: item.CreatedAt}, {Key: "updatedAt", Value: item.UpdatedAt},
        }

        interfaceArray = append(interfaceArray, airportToInterface)
        
    }

    return interfaceArray
}