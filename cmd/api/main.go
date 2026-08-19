package main

import (
	"github.com/davi1985/gopportunities/internal/config"
	"github.com/davi1985/gopportunities/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
