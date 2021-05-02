package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kodah/blog/service"
)

func LogoutRoute(group *gin.RouterGroup) {
	group.GET("/logout", func(ctx *gin.Context) {
		session := service.NewSessionService(ctx, true)

		session.Clear()

		err := session.Save()
		if err != nil {
			log.Printf("Error while saving session. error=%s", err)
		}

		// redirect users to the home page
		params := struct {
			Redirect bool `form:"redirect"`
		}{}

		err = ctx.ShouldBind(&params)
		if err != nil {
			log.Printf("Error while binding request parameters. error=%s", err)

			ctx.AbortWithStatusJSON(http.StatusBadRequest, nil)
		}

		if params.Redirect {
			ctx.Redirect(http.StatusFound, "/")

			return
		}

		ctx.JSON(http.StatusOK, nil)
	})
}
