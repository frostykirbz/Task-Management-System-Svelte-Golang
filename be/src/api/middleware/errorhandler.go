package middleware

import "github.com/gin-gonic/gin"

type Response struct {
	Status       int
	ErrorMessage string
}

func ErrorHandler(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"code": code, "message": message})
	c.Next()
}
