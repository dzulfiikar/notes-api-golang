package notes

import (
	"errors"
	"fmt"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

type CreateNoteUseCase struct {
	noteRepository      NoteRepository
	createNotePresenter CreateNotePresenter
}

func NewCreateNoteUseCase(noteRepository NoteRepository, createNotePresenter CreateNotePresenter) *CreateNoteUseCase {
	return &CreateNoteUseCase{
		noteRepository,
		createNotePresenter,
	}
}

func (useCase *CreateNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, error interface{}) {
	var createNoteDTO CreateNoteDTO
	c.BindJSON(&createNoteDTO)

	userID := c.MustGet("user_id").(string)

	if createNoteDTO.Title == "" {
		return nil, useCase.createNotePresenter.ToErrorResponse(errors.New("title is required"), 400)
	}

	if createNoteDTO.Content == "" {
		return nil, useCase.createNotePresenter.ToErrorResponse(errors.New("content is required"), 400)
	}

	result, err := useCase.noteRepository.Create(useCase.createNotePresenter.ToDomain(createNoteDTO, userID))

	if err != nil {
		return nil, useCase.createNotePresenter.ToErrorResponse(fmt.Errorf("error when creating note: %w", err), 500)
	}

	return useCase.createNotePresenter.ToResponse(result), nil
}
