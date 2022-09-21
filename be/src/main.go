package main

import (
	"backend/api/middleware"
	"backend/api/route"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	db := middleware.ConnectionToDatabase()
	defer db.Close()

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/admin-create-group", route.CreateGroup)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
