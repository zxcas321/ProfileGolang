package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	experienceService "github.com/zxcas321/ProfileGolang/services/experience"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateExperience(c *gin.Context) {
	userID, err := utils.DecodeID(c.Param("userId"))
	if err != nil {
        c.JSON(404, gin.H{"message": "invalid user Id"})
        return
    }

	var req requests.ExperienceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	experience := models.Experience{
		UserID:      userID,
		CompanyName: req.CompanyName,
		Role:        req.Role,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	if err := experienceService.CreateExperience(&experience); err != nil {
		c.JSON(500, gin.H{"message": "failed to create experience"})
		return
	}

	resp, _ := experienceService.FindByIDExperience(experience.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexExperience(c *gin.Context) {
	experience, err := experienceService.FindAllExperience()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch experience"})
		return
	}

	c.JSON(200, gin.H{"data": experience})
}

func ShowExperience(c *gin.Context) {
	experienceID, err := utils.DecodeID(c.Param("experienceId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	experience, err := experienceService.FindByIDExperience(experienceID)
	if err != nil {
		c.JSON(404, gin.H{"message": "experience not found"})
		return
	}

	c.JSON(200, gin.H{"data": experience})
}

func UpdateExperience(c *gin.Context) {
	experienceID, err := utils.DecodeID(c.Param("experienceId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
	}

	var req requests.ExperienceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	data := map[string]interface{}{}

	if req.CompanyName != "" {
		data["company_name"] = req.CompanyName
	}
	if req.Role != "" {
		data["Role"] = req.Role
	}
	data["start_date"] = req.StartDate
	data["end_date"] = req.EndDate

	if len(data) == 0 {
		c.JSON(400, gin.H{"message": "no data to update"})
		return
	}

	if err := experienceService.UpdateExperience(experienceID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	experience, _ := experienceService.FindByIDExperience(experienceID)
	c.JSON(200, gin.H{"data": experience})
}

func DeleteExperience(c *gin.Context) {
	experienceId, err := utils.DecodeID(c.Param("experienceId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	if err := experienceService.DeleteExperience(experienceId); err != nil {
		c.JSON(404, gin.H{"message": "experience not found"})
		return
	}

	c.JSON(200, gin.H{"message" : "Deleted successfully"})
}
