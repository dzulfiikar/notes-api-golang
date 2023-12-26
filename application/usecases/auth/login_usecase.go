package auth

import (
	"errors"
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/framework/sql/repositories"
	"notes-api-golang/framework/sql/schemas"

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

func (useCase *LoginUseCase) Execute(c *gin.Context) (data map[string]interface{}, err interface{}) {
	var loginDTO LoginDTO
	c.BindJSON(&loginDTO)

	user, err := useCase.userRepository.FetchUserByEmail(loginDTO.Email)

	if (err != nil && user == schemas.User{}) {
		return nil, useCase.loginPresenter.ToErrorResponse(errors.New(("Invalid Email Or Password")))
	}

	if !user.CheckPasswordMatch(loginDTO.Password) {
		return nil, useCase.loginPresenter.ToErrorResponse(errors.New(("Invalid Email Or Password")))
	}

	accessToken, _ := user.GenerateJWTToken()

	return useCase.loginPresenter.ToResponse(user, JWTToken{
		AccessToken: accessToken,
	}), nil
}
