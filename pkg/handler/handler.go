package handler

import (
	"github.com/swaggo/files"
	"github.com/gin-gonic/gin"
	"github.com/jpegShawty/go_todo_app/pkg/service"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "girhub.com/"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

// Инициализирует все наши эндпоинты
func (h *Handler) InitRoutes() *gin.Engine {
	// Инициализируем роутер
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группируем эндпоинты
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// Работа со списками и их задачами
	// Создаю группу маршрутов /api
	// Каждый путь, начинающийся с /api, будет сначала проходить через middleware h.userIdentity
	//
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			// мОЖЕТ БЫТЬ ЛЮБОЕ ЗНАЧЕНИЕ, к которому мы можем обратиться по имени параметра id
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	
	
	return router
}