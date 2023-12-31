package notes

import (
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

	result, err := useCase.noteRepository.Create(useCase.createNotePresenter.ToDomain(createNoteDTO, userID))

	if err != nil {
		return nil, useCase.createNotePresenter.ToErrorResponse(err)
	}

	return useCase.createNotePresenter.ToResponse(result), nil
}
