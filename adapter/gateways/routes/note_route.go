package routes

import (
	noteControllers "notes-api-golang/adapter/controllers/note"
	notePresenter "notes-api-golang/adapter/presenters/note"
	noteUseCases "notes-api-golang/application/note"
	"notes-api-golang/framework/http/middlewares"
	mongo "notes-api-golang/framework/mongo"
	repositories "notes-api-golang/framework/mongo/repositories"

	"github.com/gin-gonic/gin"
)

func NoteRouteGroup(router *gin.RouterGroup) {
	auth := router.Group("/notes")

	CreateNoteRoute(auth)
	FetchAllNoteRoute(auth)
	FetchNoteRoute(auth)
	DeleteNoteRoute(auth)
	UpdateNoteRoute(auth)
	RecoverNoteRoute(auth)
}

func CreateNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewCreateNotePresenter()
	useCase := noteUseCases.NewCreateNoteUseCase(*repository, presenter)
	controller := noteControllers.NewCreateNoteController(*useCase, presenter)
	router.POST("/", middlewares.CreateAuthMiddleware, controller.CreateNote)
}

func FetchAllNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewFetchAllNotePresenter()
	useCase := noteUseCases.NewFetchAllNoteUseCase(*repository, presenter)
	controller := noteControllers.NewFetchAllNoteController(*useCase, presenter)
	router.GET("/", middlewares.CreateAuthMiddleware, controller.FetchAllNote)
}

func FetchNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewFetchNotePresenter()
	useCase := noteUseCases.NewFetchNoteUseCase(*repository, presenter)
	controller := noteControllers.NewFetchNoteController(*useCase, presenter)
	router.GET("/:note_id", middlewares.CreateAuthMiddleware, controller.FetchNote)
}

func DeleteNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewDeleteNotePresenter()
	useCase := noteUseCases.NewDeleteNoteUseCase(*repository, presenter)
	controller := noteControllers.NewDeleteNoteController(*useCase, presenter)
	router.DELETE("/:note_id", middlewares.CreateAuthMiddleware, controller.DeleteNote)
}

func UpdateNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewUpdateNotePresenter()
	useCase := noteUseCases.NewUpdateNoteUseCase(*repository, presenter)
	controller := noteControllers.NewUpdateNoteController(*useCase, presenter)
	router.PUT("/:note_id", middlewares.CreateAuthMiddleware, controller.UpdateNote)
}

func RecoverNoteRoute(router *gin.RouterGroup) {
	repository := repositories.NewNoteRepository(mongo.Database)
	presenter := notePresenter.NewRecoverNotePresenter()
	useCase := noteUseCases.NewRecoverNoteUseCase(*repository, presenter)
	controller := noteControllers.NewRecoverNoteController(*useCase, presenter)
	router.PATCH("/:note_id/recover", middlewares.CreateAuthMiddleware, controller.RecoverNote)
}
