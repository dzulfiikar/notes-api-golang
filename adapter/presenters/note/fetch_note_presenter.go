package note

import (
	schema "notes-api-golang/framework/mongo/schemas"
)

type FetchNotePresenter interface {
	ToResponse(note schema.Note) (mapResponse map[string]interface{})
	ToErrorResponse(err error) (mapResponse map[string]interface{})
}

func NewFetchNotePresenter() FetchNotePresenter {
	return &fetchNotePresenter{}
}

type fetchNotePresenter struct{}

func (presenter *fetchNotePresenter) ToResponse(note schema.Note) (noteResponse map[string]interface{}) {
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

func (presenter *fetchNotePresenter) ToErrorResponse(err error) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
	}

	return
}
