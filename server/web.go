package server

import (
	"github.com/kodah/blog/service"

	"github.com/kodah/blog/api"

	"github.com/gin-gonic/gin"
)

func NewWebServer() (*gin.Engine, error) {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(service.NewCookieStore("secret_data"))
	router.Static("/static", "frontend/static")

	// register API's
	api.RegisterAPIs(router)

	return router, nil
}
