package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func NonProtectedRoutes(router *gin.RouterGroup, systemUserService services.SystemUserService) {
	controller := controllers.NewSystemUserController(systemUserService)

	routes := router.Group("/system-user")
	{
		routes.POST("/login", controller.Login)
	}
}
