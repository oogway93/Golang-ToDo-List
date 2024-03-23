package handler_todo

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todo_list/docs"
	"todo_list/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	route := gin.New()

	api := route.Group("/api")
	{
		list := api.Group("/list")
		{
			list.GET("/", h.getList)
			list.POST("/", h.createList)
			list.GET("/:id", h.getListByID)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			item := list.Group(":id/item")
			{
				item.GET("/", h.getItem)
				item.POST("/", h.createItem)
			}
		}
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return route
}
