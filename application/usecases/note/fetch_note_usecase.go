package notes

import (
	"errors"
	"fmt"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

type FetchNoteUseCase struct {
	noteRepository     NoteRepository
	fetchNotePresenter FetchNotePresenter
}

func NewFetchNoteUseCase(noteRepository NoteRepository, fetchNotePresenter FetchNotePresenter) *FetchNoteUseCase {
	return &FetchNoteUseCase{
		noteRepository,
		fetchNotePresenter,
	}
}

func (useCase *FetchNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, error interface{}) {
	noteId := c.Param("note_id")
	userID := c.MustGet("user_id").(string)

	if noteId == "" {
		return nil, useCase.fetchNotePresenter.ToErrorResponse(errors.New("note id is required"), 400)
	}

	result, err := useCase.noteRepository.FetchNoteById(noteId, userID)

	if result.ID == "" {
		return nil, useCase.fetchNotePresenter.ToErrorResponse(errors.New("note not found"), 404)
	}

	if err != nil {
		return nil, useCase.fetchNotePresenter.ToErrorResponse(fmt.Errorf("error when fetching note %w", err), 500)
	}

	return useCase.fetchNotePresenter.ToResponse(result), nil
}
