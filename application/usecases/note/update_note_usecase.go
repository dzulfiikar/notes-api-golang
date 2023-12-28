package notes

import (
	"errors"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

type UpdateNoteUseCase struct {
	noteRepository      NoteRepository
	updateNotePresenter UpdateNotePresenter
}

func NewUpdateNoteUseCase(noteRepository NoteRepository, updateNotePresenter UpdateNotePresenter) *UpdateNoteUseCase {
	return &UpdateNoteUseCase{
		noteRepository,
		updateNotePresenter,
	}
}

func (useCase *UpdateNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, err error) {
	noteId := c.Param("note_id")
	userID := c.MustGet("user_id").(string)
	var updateNoteDTO UpdateNoteDTO
	c.BindJSON(&updateNoteDTO)

	result, err := useCase.noteRepository.Update(noteId, useCase.updateNotePresenter.ToDomain(updateNoteDTO), userID)

	if result.ID == "" {
		return nil, errors.New("Note not found")
	}

	if err != nil {
		return nil, errors.New("Note not found")
	}

	return useCase.updateNotePresenter.ToResponse(result), nil
}
