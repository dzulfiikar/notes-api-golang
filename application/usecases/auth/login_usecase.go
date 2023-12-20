package auth

import (
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/framework/sql/repositories"

	"github.com/gin-gonic/gin"
)

type LoginUseCase struct {
	userRepository UserRepository
	loginPresenter LoginPresenter
}

func NewLoginUseCase(userRepository UserRepository, loginPresenter LoginPresenter) *LoginUseCase {
	return &LoginUseCase{
		userRepository,
		loginPresenter,
	}
}

func (useCase *LoginUseCase) Execute(c *gin.Context) {
	var loginDTO LoginDTO
	c.BindJSON(&loginDTO)

	user, err := useCase.userRepository.FetchUserByEmail(loginDTO.Email)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "email or password is wrong",
		})

		return
	}

	if !user.CheckPasswordMatch(loginDTO.Password) {
		c.JSON(400, gin.H{
			"message": "email or password is wrong",
		})

		return
	}

	accessToken, err := user.GenerateJWTToken()

	c.JSON(200, gin.H{
		"message": "success",
		"data": useCase.loginPresenter.ToResponse(user, JWTToken{
			AccessToken: accessToken,
		}),
	})

}
