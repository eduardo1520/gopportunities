package main

import (
	"github.com/eduardo1520/gopportunities/config"
	"github.com/eduardo1520/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
