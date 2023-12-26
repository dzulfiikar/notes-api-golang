package notes

import (
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

type FetchAllNoteUseCase struct {
	noteRepository        NoteRepository
	fetchAllNotePresenter FetchAllNotePresenter
}

func NewFetchAllNoteUseCase(noteRepository NoteRepository, fetchAllNotePresenter FetchAllNotePresenter) *FetchAllNoteUseCase {
	return &FetchAllNoteUseCase{
		noteRepository,
		fetchAllNotePresenter,
	}
}

func (useCase *FetchAllNoteUseCase) Execute(c *gin.Context) (data []map[string]interface{}, err error) {

	userID := c.MustGet("user_id").(string)

	result, err := useCase.noteRepository.FetchAllNotes(userID)

	if err != nil {
		return nil, err
	}

	return useCase.fetchAllNotePresenter.ToResponse(result), nil
}
