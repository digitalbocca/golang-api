package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func GetVehicle(c *gin.Context) {
	vehicle := Vehicle{
		ID:    "1",
		Brand: "Ford",
		Model: "Fiesta",
	}

	c.JSON(http.StatusOK, vehicle)
}

func main() {
	router := gin.Default()

	router.GET("/vehicle", GetVehicle)

	router.Run(":3000")
}
