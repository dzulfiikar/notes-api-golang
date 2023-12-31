package auth

import (
	schema "notes-api-golang/framework/sql/schemas"
)

type ProfilePresenter interface {
	ToResponse(user schema.User) (mapResponse map[string]interface{})
	ToErrorResponse(err error) (mapResponse map[string]interface{})
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

func (presenter *profilePresenter) ToErrorResponse(err error) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"message": err.Error(),
	}

	return
}
