package controllers

import (
	"net/http"

	"last/interview/models"

	"github.com/gin-gonic/gin"
)

func Fizzbuzz(c *gin.Context) {
	var ReqBody models.FizzbuzzReq

	err := c.ShouldBindJSON(&ReqBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if ReqBody.Number%3 == 0 && ReqBody.Number%5 == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error":  nil,
			"result": "FizzBuzz",
		})
		return
	}

	if ReqBody.Number%3 == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error":  nil,
			"result": "Fizz",
		})
		return
	}

	if ReqBody.Number%5 == 0 {
		c.JSON(http.StatusOK, gin.H{
			"error":  nil,
			"result": "Buzz",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":  nil,
		"result": "",
	})
}
