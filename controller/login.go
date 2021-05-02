package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/kodah/blog/dto"
	"github.com/kodah/blog/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials

	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found."
	}

	uraw, ok := ctx.Get("Username")
	if ok && uraw.(string) != "" {
		credential.Username = uraw.(string)
	}

	praw, ok := ctx.Get("Password")
	if ok && praw.(string) != "" {
		credential.Password = praw.(string)
	}

	user := controller.loginService.LoginUser(credential.Username, credential.Password)

	return controller.jWtService.GenerateToken(credential.Username, user.IsAuthenticated(), user.IsAdmin())
}
