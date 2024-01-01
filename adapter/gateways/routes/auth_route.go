package routes

import (
	controllers "notes-api-golang/adapter/controllers/auth"
	presenters "notes-api-golang/adapter/presenters/auth"
	useCases "notes-api-golang/application/auth"
	middlewares "notes-api-golang/framework/http/middlewares"
	"notes-api-golang/framework/sql"
	repository "notes-api-golang/framework/sql/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRouteGroup(router *gin.RouterGroup) {
	auth := router.Group("/auth")

	CreateSignUpRoute(auth)
	CreateLoginRoute(auth)
	CreateProfileRoute(auth)

}

func CreateSignUpRoute(router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(sql.SqlDB)
	presenter := presenters.NewSignUpPresenter()
	useCase := useCases.NewSignUpUseCase(*userRepository, presenter)
	controller := controllers.NewSignUpController(*useCase)

	router.POST("/signup", controller.SignUp)
}

func CreateLoginRoute(router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(sql.SqlDB)
	presenter := presenters.NewLoginPresenter()
	useCase := useCases.NewLoginUseCase(*userRepository, presenter)
	controller := controllers.NewLoginController(*useCase, presenter)

	router.POST("/login", controller.Login)
}

func CreateProfileRoute(router *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(sql.SqlDB)
	presenter := presenters.NewProfilePresenter()
	useCase := useCases.NewProfileUseCase(*userRepository, presenter)
	controller := controllers.NewProfileController(*useCase)

	router.GET("/profile", middlewares.CreateAuthMiddleware, controller.GetProfile)
}
