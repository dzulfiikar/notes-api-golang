package auth

import (
	schema "notes-api-golang/framework/sql/schemas"

	"github.com/google/uuid"
)

type SignUpDTO struct {
	ID       uuid.UUID
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpPresenter interface {
	ToResponse(user schema.User) (mapResponse map[string]interface{})
	ToErrorResponse(err error) (mapResponse map[string]interface{})
	ToDomain(dto SignUpDTO) (user schema.User)
}

func NewSignUpPresenter() SignUpPresenter {
	return &signUpPresenter{}
}

type signUpPresenter struct{}

func (presenter *signUpPresenter) ToResponse(user schema.User) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"id":         user.ID,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	return

}

func (presenter *signUpPresenter) ToErrorResponse(err error) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
	}
	return
}

func (presenter *signUpPresenter) ToDomain(dto SignUpDTO) (user schema.User) {
	user = schema.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	return
}
