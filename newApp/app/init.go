package app

import (
	// "touch/service/user"
	"touch/app/user"
	userService "touch/service/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/basic/docs"
)

type App struct {
	user *user.Controller
	// company *company.Controller
}

func New(userService userService.Service) *App {
	return &App{
		user: user.New(userService),
		// company: company.New(companyService),
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	docs.SwaggerInfo.Title = "Touch Tech API"
	docs.SwaggerInfo.Description = "API Spec Demo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		// 	apiRoutes.GET("/companies", app.company.Update)
		// 	apiRoutes.POST("/companies", app.company.Create)
		// 	apiRoutes.GET("/companies/:id", app.company.Read)
		// 	apiRoutes.PUT("/companies/:id", app.company.Update)
		// 	apiRoutes.DELETE("/companies/:id", app.company.Delete)

		// apiRoutes.GET("/staffs", app.user.Update)
		apiRoutes.POST("/staffs", app.user.Create)
		// apiRoutes.GET("/staffs/:id", app.user.Read)
		// apiRoutes.PUT("/staffs/:id", app.user.Update)
		// apiRoutes.DELETE("/staffs/:id", app.user.Delete)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
