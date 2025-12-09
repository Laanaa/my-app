package router

import (
	"github.com/Laanaa/my-app/internal/person"
	"github.com/gin-gonic/gin"
)


func Router() *gin.Engine {
	r := gin.Default()

	repo := person.NewRepository()
	service := person.NewService(repo)
	controller := person.NewController(service)

	api := r.Group("/persons")
	{
		api.GET("/", controller.GetAll)
		api.GET("/:id", controller.GetByID)
		api.PUT("/:id", controller.Update)
		api.POST("/create", controller.Create)
		api.DELETE("/:id", controller.Delete)
	}

	return r
}