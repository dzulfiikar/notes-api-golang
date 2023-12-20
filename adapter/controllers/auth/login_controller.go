package auth

import (
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUseCase   LoginUseCase
	loginPresenter LoginPresenter
}

func NewLoginController(loginUseCase LoginUseCase, loginPresenter LoginPresenter) *LoginController {
	return &LoginController{
		loginUseCase,
		loginPresenter,
	}
}

func (controller *LoginController) Login(c *gin.Context) {
	controller.loginUseCase.Execute(c)
}
