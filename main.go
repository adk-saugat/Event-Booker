package main

import (
	"github.com/event-booker/db"
	"github.com/event-booker/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
