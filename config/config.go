package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() error {
	// Initialize SQLite
	if db, err = InitializeSQLite(); err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize Logger
	return NewLogger(p)
}
