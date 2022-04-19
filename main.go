package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teandresmith/airplane-reservation-system/middleware"
	"github.com/teandresmith/airplane-reservation-system/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routes.PublicRoutes(router)

	router.Use(middleware.Authorization())
	routes.PrivateRoutes(router)
	

	log.Fatal(router.Run())
}