package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func ClientRoutes(router *gin.RouterGroup, service services.ClientService) {
	controller := controllers.NewClientController(service)
	routes := router.Group("/clients")
	{
		routes.POST("/create", controller.Create)
		routes.PUT("/update", controller.Update)
		routes.DELETE(":id", controller.Delete)
		routes.GET(":id", controller.GetById)
		routes.GET("/all", controller.GetAll)
	}
}
