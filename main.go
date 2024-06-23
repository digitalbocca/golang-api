package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID        uint      `json:"id"`
	Brand     string    `json:"brand"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetVehicle(c *gin.Context) {
	vehicle := Vehicle{
		Brand: "Ford",
		Model: "Fiesta",
	}

	c.JSON(http.StatusOK, vehicle)
}

func dbConn() *gorm.DB {
	/**
	 * @todo Use environment variables to set the connection string
	 */
	dsn := "host=localhost user=admin password=admin dbname=golang-api port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func migrate(db *gorm.DB, entities ...interface{}) {
	err := db.AutoMigrate(entities...)
	if err != nil {
		panic("failed to migrate")
	}
}

func main() {
	db := dbConn()

	/**
	 * @example
	 *
	 * models := []interface{}{}
	 * models = append(models, &Vehicle{})
	 * models = append(models, &User{})
	 */

	models := []interface{}{
		&Vehicle{},
		&User{},
	}

	migrate(db, models...)

	router := gin.Default()

	router.GET("/vehicle", GetVehicle)

	router.Run(":3000")
}
