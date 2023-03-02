package main

import (
	"farhadiis/order-process/infrastructure/datastore/mysql"
	"farhadiis/order-process/infrastructure/datastore/redis"
	"farhadiis/order-process/infrastructure/router"
	"farhadiis/order-process/registry"
	"farhadiis/order-process/utils"
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
	log.SetPrefix("[Orders] ")
}

func main() {
	redisClient := redis.Connect()
	defer redis.Disconnect(redisClient)

	mySqlClient := mysql.Connect()
	defer mysql.Disconnect(mySqlClient)

	reg := registry.NewRegistry(redisClient, mySqlClient)
	app := router.NewRouter(gin.Default(), reg.NewAppController())
	app.SetTrustedProxies(nil)

	port := utils.GetEnv("PORT")
	log.Printf("Server listen at port %s", port)
	err := app.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
