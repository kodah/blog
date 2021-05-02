package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kodah/blog/api/v1"
)

func RegisterAPIs(router *gin.Engine) {
	api := router.Group("/api")
	v1.RegisterGroup(api)
}
