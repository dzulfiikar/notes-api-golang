package auth

import (
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/framework/sql/repositories"

	"github.com/gin-gonic/gin"
)

type ProfileUseCase struct {
	userRepository   UserRepository
	profilePresenter ProfilePresenter
}

// crate new instance of ProfileUseCase
func NewProfileUseCase(userRepository UserRepository, profilePresenter ProfilePresenter) *ProfileUseCase {
	return &ProfileUseCase{
		userRepository,
		profilePresenter,
	}
}

// execute ProfileUseCase
func (useCase *ProfileUseCase) Execute(c *gin.Context) (data map[string]interface{}, err error) {
	var userId = c.MustGet("user_id").(string)

	user, err := useCase.userRepository.FetchUserById(userId)
	if err != nil {
		c.JSON(403, gin.H{
			"message": "error",
		})

		return nil, err
	}

	return useCase.profilePresenter.ToResponse(user), nil
}
