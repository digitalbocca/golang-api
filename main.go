package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/digitalbocca/golang-api/orm"
	"github.com/digitalbocca/golang-api/structs"
)

func GetVehicle(c *gin.Context) {
	vehicle := structs.Vehicle{
		Brand: "Ford",
		Model: "Fiesta",
	}

	c.JSON(http.StatusOK, vehicle)
}

func main() {
	db := orm.DbConn()

	/**
	 * @example Using variadic function to pass multiple models
	 *
	 * models := []interface{}{}
	 * models = append(models, &Vehicle{})
	 * models = append(models, &User{})
	 */

	models := []interface{}{
		&structs.Vehicle{},
		&structs.User{},
	}

	orm.Migrate(db, models...)

	router := gin.Default()

	router.GET("/vehicle", GetVehicle)

	router.Run(":3000")
}
