package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type RecoverNoteController struct {
	recoverNoteController notesUseCase.RecoverNoteUseCase
	recoverNotePresenter  notesPresenter.RecoverNotePresenter
}

func NewRecoverNoteController(recoverNoteUseCase notesUseCase.RecoverNoteUseCase, recoverNotePresenter notesPresenter.RecoverNotePresenter) *RecoverNoteController {
	return &RecoverNoteController{
		recoverNoteUseCase,
		recoverNotePresenter}
}

func (controller *RecoverNoteController) RecoverNote(c *gin.Context) {
	result, err := controller.recoverNoteController.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return
	}

	responses.NewSuccessResponse(result).Send(c)

}
