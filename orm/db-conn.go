package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConn() *gorm.DB {
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
