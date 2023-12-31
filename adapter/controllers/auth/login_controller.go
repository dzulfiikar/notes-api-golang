package auth

import (
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/application/usecases/auth"
	"notes-api-golang/framework/http/responses"

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
	result, err := controller.loginUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
