package note

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/mongo/schemas"
)

type RecoverNotePresenter interface {
	Presenter
	ToResponse(note schema.Note) (mapResponse map[string]interface{})
}

func NewRecoverNotePresenter() RecoverNotePresenter {
	return &recoverNotePresenter{}
}

type recoverNotePresenter struct{}

func (presenter *recoverNotePresenter) ToResponse(note schema.Note) (noteResponse map[string]interface{}) {
	noteResponse = map[string]interface{}{
		"id":         note.ID,
		"title":      note.Title,
		"content":    note.Content,
		"created_at": note.CreatedAt,
		"created_by": note.CreatedBy,
		"updated_at": note.UpdatedAt,
		"updated_by": note.UpdatedBy,
		"deleted":    note.Deleted,
		"deleted_at": note.DeletedAt,
		"deleted_by": note.DeletedBy,
	}

	return

}

func (presenter *recoverNotePresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
		"code":  code,
	}

	return
}
