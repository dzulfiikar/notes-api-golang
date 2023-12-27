package note

import (
	schema "notes-api-golang/framework/mongo/schemas"
)

type DeleteNotePresenter interface {
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
