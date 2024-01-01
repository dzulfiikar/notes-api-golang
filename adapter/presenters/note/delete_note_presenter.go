package note

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/mongo/schemas"
)

type DeleteNotePresenter interface {
	Presenter
	ToResponse(note schema.Note) (mapResponse map[string]interface{})
}

func NewDeleteNotePresenter() DeleteNotePresenter {
	return &deleteNotePresenter{}
}

type deleteNotePresenter struct{}

func (presenter *deleteNotePresenter) ToResponse(note schema.Note) (noteResponse map[string]interface{}) {
	noteResponse = map[string]interface{}{
		"id": note.ID,
	}

	return

}

func (presenter *deleteNotePresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
		"code":  code,
	}
	return
}
