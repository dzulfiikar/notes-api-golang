package auth

import (
	. "notes-api-golang/application/usecases/auth"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileUseCase ProfileUseCase
}

func NewProfileController(profileUseCase ProfileUseCase) *ProfileController {
	return &ProfileController{
		profileUseCase,
	}
}

func (controller *ProfileController) GetProfile(c *gin.Context) {
	result, err := controller.profileUseCase.Execute(c)
	if err != nil {
		responses.NewBadRequestError(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
