package config

import (
	"os"

	"github.com/eduardo1520/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if DB Exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		// Create DB
		logger.Info("sqlite db not found, creating...")

		// Create the database file and directory
		if err = os.MkdirAll("./db", os.ModePerm); err != nil {
			logger.Errorf("sqlite db creation error: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("sqlite db creation error: %v", err)
			return nil, err
		}

		// Close the file
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil

}
