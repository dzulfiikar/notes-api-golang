package auth

import (
	. "notes-api-golang/application/usecases/auth"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	signUpUseCase SignUpUseCase
}

func NewSignUpController(signUpUseCase SignUpUseCase) *SignUpController {
	return &SignUpController{
		signUpUseCase: signUpUseCase,
	}
}

func (controller *SignUpController) SignUp(c *gin.Context) {
	controller.signUpUseCase.Execute(c)
}
