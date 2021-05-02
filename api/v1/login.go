package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kodah/blog/controller"
	"github.com/kodah/blog/service"
)

func LoginRoute(group *gin.RouterGroup) {
	group.POST("/login", func(ctx *gin.Context) {
		var configService service.ConfigService = service.ConfigurationService("")
		if configService.Error() != nil {
			ctx.AbortWithError(http.StatusInternalServerError, configService.Error())

			return
		}

		var dbService service.DBService = service.SQLiteDBService("")
		if dbService.Error() != nil {
			ctx.AbortWithError(http.StatusInternalServerError, dbService.Error())

			return
		}

		var loginService service.LoginService = service.DynamicLoginService(dbService)
		var jwtService service.JWTService = service.JWTAuthService()
		var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

		session := service.NewSessionService(ctx, false)

		var data gin.H
		token := loginController.Login(ctx)
		status := http.StatusInternalServerError

		if token != "" {
			session.Set("token", token)

			status = http.StatusOK
			data = gin.H{
				"token": token,
			}
		} else {
			session.Set("token", "")
		}

		err := session.Save()
		if err != nil {
			log.Printf("Error while saving session. error=%s", err)
		}

		ctx.JSON(status, data)
	})
}
