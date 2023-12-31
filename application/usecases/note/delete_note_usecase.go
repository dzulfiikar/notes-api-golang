package notes

import (
	"errors"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

type DeleteNoteUseCase struct {
	noteRepository      NoteRepository
	deleteNotePresenter DeleteNotePresenter
}

func NewDeleteNoteUseCase(noteRepository NoteRepository, deleteNotePresenter DeleteNotePresenter) *DeleteNoteUseCase {
	return &DeleteNoteUseCase{
		noteRepository,
		deleteNotePresenter,
	}
}

func (useCase *DeleteNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, err interface{}) {
	noteId := c.Param("note_id")
	userID := c.MustGet("user_id").(string)

	result, err := useCase.noteRepository.SoftDelete(noteId, userID)

	if result.ID == "" {
		return nil, useCase.deleteNotePresenter.ToErrorResponse(errors.New("note not found"))
	}

	if err != nil {
		return nil, useCase.deleteNotePresenter.ToErrorResponse(errors.New("note not found"))
	}

	return useCase.deleteNotePresenter.ToResponse(result), nil
}
