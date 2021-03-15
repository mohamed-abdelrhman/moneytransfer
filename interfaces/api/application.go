package api

import (
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/moneytransfer/infrastructure/clients"
	"log"
)

var(
	Router=gin.Default()
)

func StartApplication() {
	clients.GetRedisClient()
	CustomerUrlMapping()
	TransferUrlMapping()
	Router.Use(cors.AllowAll())
	log.Fatal(Router.Run(":8080"))
}
