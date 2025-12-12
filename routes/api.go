package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/controllers/api"
)

func ApiRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("", api.IndexUser)
			userGroup.GET("/:id", api.ShowUser)
			userGroup.POST("", api.CreateUser)
			userGroup.PUT("/:id", api.UpdateUser)
			userGroup.DELETE("/:id", api.DeleteUser)
			
			projectGroup := userGroup.Group("/:userId/project")
			{
				projectGroup.GET("", api.IndexProject)
				projectGroup.GET("/:projectId", api.ShowProject)
				projectGroup.POST("", api.CreateProject)
				projectGroup.PUT("/:projectId", api.UpdateProject)
				projectGroup.DELETE("/:projectId", api.DeleteProject)

   				skill := projectGroup.Group("/:projectId/skill")
				{
					skill.GET("", api.IndexSkill)
					skill.GET("/:skillId", api.ShowSkill)
					skill.POST("", api.CreateSkill)
					skill.PUT("/:skillId", api.UpdateSkill)
					skill.DELETE("/:skillId", api.DeleteSkill)
				}
			}



		}
	}
}
