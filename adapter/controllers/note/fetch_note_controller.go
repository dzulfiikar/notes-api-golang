package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type FetchNoteController struct {
	fetchNoteUseCase   notesUseCase.FetchNoteUseCase
	fetchNotePresenter notesPresenter.FetchNotePresenter
}

func NewFetchNoteController(fetchNoteUseCase notesUseCase.FetchNoteUseCase, fetchNotePresenter notesPresenter.FetchNotePresenter) *FetchNoteController {
	return &FetchNoteController{
		fetchNoteUseCase,
		fetchNotePresenter,
	}
}

func (controller *FetchNoteController) FetchNote(c *gin.Context) {
	result, err := controller.fetchNoteUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
