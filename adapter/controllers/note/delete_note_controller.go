package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/usecases/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type DeleteNoteController struct {
	deleteNoteUseCase   notesUseCase.DeleteNoteUseCase
	deleteNotePresenter notesPresenter.DeleteNotePresenter
}

func NewDeleteNoteController(deleteNoteUseCase notesUseCase.DeleteNoteUseCase, deleteNotePresenter notesPresenter.DeleteNotePresenter) *DeleteNoteController {
	return &DeleteNoteController{
		deleteNoteUseCase,
		deleteNotePresenter,
	}
}

func (controller *DeleteNoteController) DeleteNote(c *gin.Context) {
	result, err := controller.deleteNoteUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
