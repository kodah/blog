package v1

import "github.com/gin-gonic/gin"

func RegisterGroup(router *gin.RouterGroup) {
	v1 := router.Group("v1")

	LoginRoute(v1)
	LogoutRoute(v1)
}
