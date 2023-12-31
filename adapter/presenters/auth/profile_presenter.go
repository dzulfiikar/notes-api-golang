package auth

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/sql/schemas"
)

type ProfilePresenter interface {
	Presenter
	ToResponse(user schema.User) (mapResponse map[string]interface{})
}

func NewProfilePresenter() ProfilePresenter {
	return &profilePresenter{}
}

type profilePresenter struct{}

func (presenter *profilePresenter) ToResponse(user schema.User) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"user": map[string]interface{}{"id": user.ID, "email": user.Email, "created_at": user.CreatedAt, "updated_at": user.UpdatedAt},
	}

	return

}

func (presenter *profilePresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"message": err.Error(),
		"code":    code,
	}

	return
}
