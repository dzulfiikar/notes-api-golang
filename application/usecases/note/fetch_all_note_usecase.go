package notes

import (
	"fmt"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func (useCase *FetchAllNoteUseCase) Execute(c *gin.Context) (data []map[string]interface{}, error interface{}) {

	userID := c.MustGet("user_id").(string)

	filter := bson.M{"created_by": userID}

	// if query deleted is true, then fetch all notes including deleted notes
	if c.Query("deleted") == "true" {
		filter = bson.M{"created_by": userID, "deleted": true}
	} else {
		filter = bson.M{"created_by": userID, "deleted": bson.M{"$exists": false}}
	}

	result, err := useCase.noteRepository.FetchAllNotes(filter)

	if err != nil {
		return nil, useCase.fetchAllNotePresenter.ToErrorResponse(fmt.Errorf("error when fething all note %w", err), 500)
	}

	return useCase.fetchAllNotePresenter.ToResponse(result), nil
}
