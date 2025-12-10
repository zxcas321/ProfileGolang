package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateExperience(c *gin.Context) {
	var req requests.ExperienceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	experience := models.Experience{
		CompanyName: req.CompanyName,
		Role:        req.Role,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	if err := services.CreateExperience(&experience); err != nil {
		c.JSON(500, gin.H{"message": "failed to create experience"})
		return
	}

	resp, _ := services.FindByIDExperience(experience.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexExperience(c *gin.Context) {
	experience, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch experience"})
	}

	c.JSON(200, gin.H{"data": experience})
}

func ShowExperience(c *gin.Context) {
	experienceID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	experience, err := services.FindByIDExperience(experienceID)
	if err != nil {
		c.JSON(404, gin.H{"message": "experience not found"})
		return
	}

	c.JSON(200, gin.H{"data": experience})
}

func UpdateExperience(c *gin.Context) {
	experienceID, err := utils.DecodeID(c.Param("id"))
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

	if err := services.UpdateExperience(experienceID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	experience, _ := services.FindByIDExperience(experienceID)
	c.JSON(200, gin.H{"data": experience})
}
