package note

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/mongo/schemas"
)

type FetchAllNotePresenter interface {
	Presenter
	ToResponse(notes []schema.Note) (mapResponse []map[string]interface{})
}

func NewFetchAllNotePresenter() FetchAllNotePresenter {
	return &fetchAllNotePresenter{}
}

type fetchAllNotePresenter struct{}

func (presenter *fetchAllNotePresenter) ToResponse(notes []schema.Note) (notesResponse []map[string]interface{}) {
	for _, note := range notes {
		notesResponse = append(notesResponse, map[string]interface{}{
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
		})
	}

	return

}

func (presenter *fetchAllNotePresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
		"code":  code,
	}

	return
}
