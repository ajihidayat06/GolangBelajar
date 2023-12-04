package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbUrl := fmt.Sprintf(`postgres://%s:%s@localhost:%s/crud`, "root", "secret", "5432")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateModel(db *gorm.DB, model ...interface{}) {
	db.AutoMigrate(model...)
}
