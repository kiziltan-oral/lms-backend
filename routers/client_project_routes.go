package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func ClientProjectRoutes(router *gin.RouterGroup, service services.ClientProjectService) {
	controller := controllers.NewClientProjectController(service)
	routes := router.Group("/client-projects")
	{
		routes.POST("/create", controller.Create)
		routes.PUT("/update", controller.Update)
		routes.DELETE(":id", controller.Delete)
		routes.GET(":id", controller.GetById)
		routes.GET("/all", controller.GetAll)
		routes.GET("/client/:clientId", controller.GetByClientId)
	}
}
