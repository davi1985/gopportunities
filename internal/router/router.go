package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Init Router
	router := gin.Default()

	// Init routes
	initializeRoutes(router)

	// Run the server
	router.Run()
}
