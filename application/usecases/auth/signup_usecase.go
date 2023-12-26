package auth

import (
	"errors"
	. "notes-api-golang/adapter/presenters/auth"
	. "notes-api-golang/framework/sql/repositories"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpUseCase struct {
	userRepository  UserRepository
	signUpPresenter SignUpPresenter
}

// crate new instance of SignUpUseCase
func NewSignUpUseCase(userRepository UserRepository, signUpPresenter SignUpPresenter) *SignUpUseCase {
	return &SignUpUseCase{
		userRepository,
		signUpPresenter,
	}
}

// execute SignUpUseCase
func (useCase *SignUpUseCase) Execute(c *gin.Context) (data map[string]interface{}, err interface{}) {
	var signUpDto SignUpDTO
	c.BindJSON(&signUpDto)

	// check if email already exist
	if useCase.userRepository.FetchUserExistsByEmail(signUpDto.Email) {
		return nil, useCase.signUpPresenter.ToBadRequestResponse(errors.New("Email already exists"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	signUpDto.Password = string(hashedPassword)

	// save user
	var user = useCase.userRepository.Save(useCase.signUpPresenter.ToDomain(signUpDto))

	return useCase.signUpPresenter.ToResponse(user), nil

}
