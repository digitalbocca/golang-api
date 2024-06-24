package orm

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, entities ...interface{}) {
	err := db.AutoMigrate(entities...)
	if err != nil {
		panic("failed to migrate")
	}
}
