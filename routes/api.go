package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/controllers/api"
)

func ApiRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		users := apiGroup.Group("/users")
		{
			users.GET("", api.IndexUser)
			users.GET("/:id", api.ShowUser)
			users.POST("", api.CreateUser)
			users.PUT("/:id", api.UpdateUser)
			users.DELETE("/:id", api.DeleteUser)
		}
	}
}
