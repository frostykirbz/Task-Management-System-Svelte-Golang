package main

import (
	"backend/api/middleware"
	"backend/api/route"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := middleware.ConnectionToDatabase()
	defer db.Close()

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/admin-create-group", route.AdminCreateGroup)

	port := middleware.LoadENV("SERVER_PORT")
	server := fmt.Sprintf(":%v", port)

	router.Run(server)
}
