package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	skillService "github.com/zxcas321/ProfileGolang/services/project"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateSkill(c *gin.Context) {
	projectID, err := utils.DecodeID(c.Param("projectId"))
	if err != nil {
        c.JSON(404, gin.H{"message": "invalid project Id"})
        return
    }

	var req requests.SkillRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	skill := models.Skills{
		ProjectID : projectID,
		SkillName: req.SkillName,
		Category:  req.Category,
	}

	if err := skillService.CreateSkill(&skill); err != nil {
		c.JSON(500, gin.H{"message": "failed to create skill"})
		return
	}

	resp, _ := skillService.FindByIDSkill(skill.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexSkill(c *gin.Context) {
	skill, err := skillService.FindAllSkill()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch skill"})
		return
	}

	c.JSON(200, gin.H{"data": skill})
}

func ShowSkill(c *gin.Context) {
	skillID, err := utils.DecodeID(c.Param("skillId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	skill, err := skillService.FindByIDSkill(skillID)
	if err != nil {
		c.JSON(4040, gin.H{"message": "skill not found"})
		return
	}

	c.JSON(200, gin.H{"data": skill})
}

func UpdateSkill(c *gin.Context) {
	skillID, err := utils.DecodeID(c.Param("skillId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	var req requests.SkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	data := map[string]interface{}{}

	if req.SkillName != "" {
		data["skill_name"] = req.SkillName
	}
	if req.Category != "" {
		data["category"] = req.Category
	}

	if len(data) == 0 {
		c.JSON(400, gin.H{"message": "no data to update"})
		return
	}

	if err := skillService.UpdateSkill(skillID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	skill, _ := skillService.FindByIDSkill(skillID)
	c.JSON(200, gin.H{"data": skill})
}

func DeleteSkill(c *gin.Context) {
	SkillID, err := utils.DecodeID(c.Param("skillId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	if err := skillService.DeleteSkill(SkillID); err != nil {
		c.JSON(404, gin.H{"message": "skill not found"})
		return
	}

	c.JSON(200, gin.H{"message" : "skill deleted successfully"})
}
