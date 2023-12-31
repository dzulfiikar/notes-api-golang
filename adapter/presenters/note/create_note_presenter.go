package note

import (
	. "notes-api-golang/adapter/presenters"
	schema "notes-api-golang/framework/mongo/schemas"
	"time"
)

type CreateNoteDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateNotePresenter interface {
	Presenter
	ToResponse(note schema.Note) (mapResponse map[string]interface{})
	ToDomain(dto CreateNoteDTO, userId string) (note schema.Note)
}

func NewCreateNotePresenter() CreateNotePresenter {
	return &createNotePresenter{}
}

type createNotePresenter struct{}

func (presenter *createNotePresenter) ToResponse(note schema.Note) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"id":         note.ID,
		"title":      note.Title,
		"content":    note.Content,
		"created_at": note.CreatedAt,
		"created_by": note.CreatedBy,
		"updated_at": note.UpdatedAt,
		"updated_by": note.UpdatedBy,
		"deleted":    note.Deleted,
	}

	return

}

func (presenter *createNotePresenter) ToDomain(dto CreateNoteDTO, userId string) (note schema.Note) {
	note = schema.Note{
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: time.Now(),
		CreatedBy: userId,
		UpdatedAt: time.Now(),
		UpdatedBy: userId,
	}

	return
}

func (presenter *createNotePresenter) ToErrorResponse(err error, code int) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
	}

	return
}
