package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func SystemUserRoutes(router *gin.RouterGroup, service services.SystemUserService) {
	controller := controllers.NewSystemUserController(service)
	routes := router.Group("/system-user")
	{
		routes.POST("/create", controller.Create)
		routes.PUT("/update", controller.Update)
		routes.DELETE("/:id", controller.Delete)
		routes.GET("/:id", controller.GetById)
		routes.GET("/email", controller.GetByEmail)
		routes.GET("/all", controller.GetAll)
		routes.POST("/logout", controller.Logout)
	}
}
