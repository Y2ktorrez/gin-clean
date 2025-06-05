package v1

import (
	"gin-clean/internal/usecase/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userService services.UserService) {
	v1 := router.Group("/api/v1")
	{
		userHandler := NewUserHandler(userService)
		users := v1.Group("/users")
		{
			users.POST("/", userHandler.Create)
			users.GET("/:id", userHandler.GetByID)
			users.GET("/", userHandler.GetAll)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}
}
