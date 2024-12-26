package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func SystemUserSettingRoutes(router *gin.RouterGroup, service services.SystemUserSettingService) {
	controller := controllers.NewSystemUserSettingController(service)
	routes := router.Group("/system-user-settings")
	{
		routes.GET("/user/:userId", controller.GetByUserId)
		routes.GET("/:id", controller.GetById)
		routes.POST("/", controller.Set)
		routes.DELETE("/:id", controller.Delete)
		routes.GET("/value/:userId", controller.GetValue)
	}
}
