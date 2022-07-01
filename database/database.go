package database

import (
	"os"

	"github.com/sirupsen/logrus"
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
		dsn := os.Getenv("DATABASE_URL")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		// On error return inmediately
		if err != nil {
			return nil, err
		}

		logrus.Info("Database initialized successfully")
	}
	// On success, return database object
	return db, nil
}
