package auth

import (
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
func (useCase *SignUpUseCase) Execute(c *gin.Context) (err error) {
	var signUpDto SignUpDTO
	c.BindJSON(&signUpDto)

	// check if email already exist
	if useCase.userRepository.FetchUserExistsByEmail(signUpDto.Email) {
		c.JSON(400, gin.H{
			"message": "email already exist",
		})

		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDto.Password), bcrypt.DefaultCost)
	if err != nil {
		// handle error
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
	}

	signUpDto.Password = string(hashedPassword)

	// save user
	var user = useCase.userRepository.Save(useCase.signUpPresenter.ToDomain(signUpDto))

	c.JSON(200, gin.H{
		"message": "success",
		"data":    useCase.signUpPresenter.ToResponse(user),
	})

	return
}
