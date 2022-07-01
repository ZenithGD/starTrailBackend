package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Single instance of the database
var db *gorm.DB = nil

// Get the single database instance. If the database hasn't been initialized yet,
// it will load the credentials from the project's enviroment variables.
// Otherwise, it will return an error.
func GetDB() (*gorm.DB, error) {
	var err error
	if db == nil {
		// Initialize database if it hasn't been yet initialized
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PORT"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		// On error return inmediately
		if err != nil {
			return nil, err
		}
	}

	// On success, return database object
	return db, nil
}
