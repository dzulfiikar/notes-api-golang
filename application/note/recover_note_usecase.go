package notes

import (
	"errors"
	"fmt"
	. "notes-api-golang/adapter/presenters/note"
	. "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecoverNoteUseCase struct {
	noteRepository       NoteRepository
	recoverNotePresenter RecoverNotePresenter
}

func NewRecoverNoteUseCase(noteRepository NoteRepository, recoverNotePresenter RecoverNotePresenter) *RecoverNoteUseCase {
	return &RecoverNoteUseCase{
		noteRepository,
		recoverNotePresenter,
	}
}

func (useCase *RecoverNoteUseCase) Execute(c *gin.Context) (data map[string]interface{}, err interface{}) {
	noteIdParam := c.Param("note_id")

	noteId, err := primitive.ObjectIDFromHex(noteIdParam)

	if err != nil {
		return nil, useCase.recoverNotePresenter.ToErrorResponse(fmt.Errorf("invalid note id: %w", err), 400)
	}

	userID := c.MustGet("user_id").(string)

	filter := bson.M{"created_by": userID, "deleted": true, "_id": noteId}
	fetchedNote, err := useCase.noteRepository.FetchAllNotes(filter)

	if len(fetchedNote) == 0 {
		return nil, useCase.recoverNotePresenter.ToErrorResponse(errors.New("note not found"), 404)
	}

	if !fetchedNote[0].Deleted {
		return nil, useCase.recoverNotePresenter.ToErrorResponse(errors.New("note is not deleted"), 400)
	}

	if err != nil {
		return nil, useCase.recoverNotePresenter.ToErrorResponse(fmt.Errorf("error when recovering note: %w", err), 500)
	}

	result, err := useCase.noteRepository.RecoverNote(noteIdParam)

	if err != nil {
		return nil, useCase.recoverNotePresenter.ToErrorResponse(fmt.Errorf("error when recovering note: %w", err), 500)
	}

	return useCase.recoverNotePresenter.ToResponse(result), nil
}
