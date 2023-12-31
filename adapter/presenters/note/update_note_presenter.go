package note

import (
	schema "notes-api-golang/framework/mongo/schemas"
)

type UpdateNoteDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateNotePresenter interface {
	ToDomain(dto UpdateNoteDTO) (note schema.Note)
	ToResponse(note schema.Note) (mapResponse map[string]interface{})
	ToErrorResponse(err error) (mapResponse map[string]interface{})
}

func NewUpdateNotePresenter() UpdateNotePresenter {
	return &updateNotePresenter{}
}

type updateNotePresenter struct{}

func (presenter *updateNotePresenter) ToResponse(note schema.Note) (noteResponse map[string]interface{}) {
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

func (presenter *updateNotePresenter) ToDomain(dto UpdateNoteDTO) (note schema.Note) {
	note = schema.Note{
		Title:   dto.Title,
		Content: dto.Content,
	}

	return
}

func (presenter *updateNotePresenter) ToErrorResponse(err error) (mapResponse map[string]interface{}) {
	mapResponse = map[string]interface{}{
		"error": err.Error(),
	}

	return
}
