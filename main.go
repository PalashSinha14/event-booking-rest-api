package main

//import "fmt"

import (
	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/routes"
	"github.com/palashsinha14/go-rest-api/db"
)

func main() {
	//fmt.Println("Hello World")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
