package auth

import (
	. "notes-api-golang/application/auth"
	"notes-api-golang/framework/http/responses"

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
	result, err := controller.signUpUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)
}
