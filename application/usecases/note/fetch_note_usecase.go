package notes

import (
	"errors"
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

func (useCase *FetchNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, err error) {
	noteId := c.Param("note_id")
	userID := c.MustGet("user_id").(string)

	result, err := useCase.noteRepository.FetchNoteById(noteId, userID)

	if result.ID == "" {
		return nil, errors.New("Note not found")
	}

	if err != nil {
		return nil, errors.New("Note not found")
	}

	return useCase.fetchNotePresenter.ToResponse(result), nil
}
