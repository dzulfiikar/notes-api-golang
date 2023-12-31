package notes

import (
	notesPresenter "notes-api-golang/adapter/presenters/note"
	notesUseCase "notes-api-golang/application/usecases/note"
	"notes-api-golang/framework/http/responses"

	"github.com/gin-gonic/gin"
)

type CreateNoteController struct {
	createNoteUseCase   notesUseCase.CreateNoteUseCase
	createNotePresenter notesPresenter.CreateNotePresenter
}

func NewCreateNoteController(createNoteUseCase notesUseCase.CreateNoteUseCase, createNotePresenter notesPresenter.CreateNotePresenter) *CreateNoteController {
	return &CreateNoteController{
		createNoteUseCase,
		createNotePresenter,
	}
}

func (controller *CreateNoteController) CreateNote(c *gin.Context) {
	result, err := controller.createNoteUseCase.Execute(c)
	if err != nil {
		responses.NewErrorResponse(err).Send(c)
		return

	}

	responses.NewErrorResponse(result).Send(c)

}
