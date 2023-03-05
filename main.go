package main

import (
	"last/interview/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/fizzbuzz", controllers.Fizzbuzz)

	r.Run()
}
