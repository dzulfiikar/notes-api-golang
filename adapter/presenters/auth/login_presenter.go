package auth

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/sql/schemas"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTToken struct {
	AccessToken string
}

type LoginPresenter interface {
	Presenter
	ToResponse(user schema.User, jwtToken JWTToken) (mapResponse map[string]interface{})
	ToDomain(dto LoginDTO) (user LoginDTO)
}

func NewLoginPresenter() LoginPresenter {
	return &loginPresenter{}
}

type loginPresenter struct{}

func (presenter *loginPresenter) ToResponse(user schema.User, jwtToken JWTToken) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"user":  map[string]interface{}{"id": user.ID, "email": user.Email, "created_at": user.CreatedAt, "updated_at": user.UpdatedAt},
		"token": map[string]interface{}{"access_token": jwtToken.AccessToken},
	}

	return

}

func (presenter *loginPresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
		"code":  code,
	}

	return
}

func (presenter *loginPresenter) ToDomain(dto LoginDTO) (user LoginDTO) {
	user = LoginDTO{
		Email:    dto.Email,
		Password: dto.Password,
	}

	return
}
