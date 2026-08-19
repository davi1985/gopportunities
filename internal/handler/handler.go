// Package handler contains HTTP handlers for the application.
package handler

import (
	"github.com/davi1985/gopportunities/internal/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
