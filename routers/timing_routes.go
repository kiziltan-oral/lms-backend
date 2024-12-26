package routers

import (
	"lms-web-services-main/controllers"
	"lms-web-services-main/services"

	"github.com/gin-gonic/gin"
)

func TimingRoutes(router *gin.RouterGroup, service services.TimingService) {
	controller := controllers.NewTimingController(service)
	routes := router.Group("/timings")
	{
		routes.POST("/create", controller.Create)
		routes.PUT("/update", controller.Update)
		routes.DELETE(":id", controller.Delete)
		routes.GET(":id", controller.GetById)
		routes.GET("/all", controller.GetAll)
		routes.GET("/client-project/:clientProjectId", controller.GetByClientProjectId)
		routes.GET("/date-range", controller.GetByDateRange)
	}
}
