package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/controllers/api"
)

func ApiRoutes(r *gin.Engine) {
	
	admin := r.Group("/admins")
	{
		admin.GET("", api.IndexAdmin)
		admin.POST("", api.CreateAdmin)
	}

	apiGroup := r.Group("/api")
	{
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("", api.IndexUser)
			userGroup.GET("/:userId", api.ShowUser)
			userGroup.POST("", api.CreateUser)
			userGroup.PUT("/:userId", api.UpdateUser)
			userGroup.DELETE("/:userId", api.DeleteUser)
			
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
			
			experience := userGroup.Group("/:userId/experience") 
			{
				experience.GET("", api.IndexExperience)
				experience.GET("/:experienceId", api.ShowExperience)
				experience.POST("", api.CreateExperience)
				experience.PUT("/:experienceId", api.UpdateExperience)
				experience.DELETE("/:experienceId", api.DeleteExperience)
			}

			certification := userGroup.Group("/:userId/certification")
			{
				certification.GET("", api.IndexCertification)
				certification.GET("/:certificationId", api.ShowCertification)
				certification.POST("", api.CreateCertification)
				certification.PUT("/:certificationId", api.UpdateCertification)
				certification.DELETE("/certificationId", api.DeleteCertification)
			}
			academic := userGroup.Group("/:userId/academic")
			{
				academic.GET("", api.IndexAcademic)
				academic.GET("/:academicId", api.ShowAcademic)
				academic.POST("", api.CreateAcademic)
				academic.PUT("/:academicId", api.UpdateAcademic)
				academic.DELETE("/academicId", api.DeleteAcademic)
			}
		}
	}
}
