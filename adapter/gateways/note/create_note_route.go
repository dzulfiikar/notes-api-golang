package note

import (
	noteControllers "notes-api-golang/adapter/controllers/note"
	notePresenter "notes-api-golang/adapter/presenters/note"
	noteUseCases "notes-api-golang/application/usecases/note"
	"notes-api-golang/framework/http/middlewares"
	mongo "notes-api-golang/framework/mongo"
	repositories "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

func NoteRouteGroup(router *gin.RouterGroup) {
	auth := router.Group("/notes")

	CreateNoteRoute(auth)
}

func CreateNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewCreateNotePresenter()
	useCase := noteUseCases.NewCreateNoteUseCase(*repository, presenter)
	controller := noteControllers.NewCreateNoteController(*useCase, presenter)
	router.POST("/", middlewares.CreateAuthMiddleware, controller.CreateNote)
}
