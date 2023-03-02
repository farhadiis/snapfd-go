package main

import (
	"farhadiis/order/infrastructure/datastore"
	"farhadiis/order/infrastructure/router"
	"farhadiis/order/registry"
	"farhadiis/order/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mode := utils.GetEnv("MODE")
	if mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	log.SetPrefix("[Order] ")
}

func main() {
	redisClient := datastore.Connect()
	defer datastore.Disconnect(redisClient)

	reg := registry.NewRegistry(redisClient)
	app := router.NewRouter(gin.Default(), reg.NewAppController())
	app.SetTrustedProxies(nil)

	port := utils.GetEnv("PORT")
	log.Printf("Server listen at port %s", port)
	err := app.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
