package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type FetchAllNoteController struct {
	fetchAllNoteUseCase   notesUseCase.FetchAllNoteUseCase
	fetchAllNotePresenter notesPresenter.FetchAllNotePresenter
}

func NewFetchAllNoteController(fetchAllNoteUseCase notesUseCase.FetchAllNoteUseCase, fetchAllNotePresenter notesPresenter.FetchAllNotePresenter) *FetchAllNoteController {
	return &FetchAllNoteController{
		fetchAllNoteUseCase,
		fetchAllNotePresenter,
	}
}

func (controller *FetchAllNoteController) FetchAllNote(c *gin.Context) {
	result, err := controller.fetchAllNoteUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
