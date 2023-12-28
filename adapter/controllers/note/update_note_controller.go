package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/usecases/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type UpdateNoteController struct {
	updateNoteUseCase   notesUseCase.UpdateNoteUseCase
	updateNotePresenter notesPresenter.UpdateNotePresenter
}

func NewUpdateNoteController(updateNoteUseCase notesUseCase.UpdateNoteUseCase, updateNotePresenter notesPresenter.UpdateNotePresenter) *UpdateNoteController {
	return &UpdateNoteController{
		updateNoteUseCase,
		updateNotePresenter,
	}
}

func (controller *UpdateNoteController) UpdateNote(c *gin.Context) {
	result, err := controller.updateNoteUseCase.Execute(c)
	if err != nil {
		responses.NewBadRequestError(err).Send(c)
		return

	}

	responses.NewSuccessResponse(result).Send(c)

}
