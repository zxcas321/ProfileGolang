package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateProject(c *gin.Context){
	var req requests.ProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message" : err.Error()})
	}

	project := models.Project{
		Name: req.Name,
		Description: req.Description,
		WebUrl: req.WebUrl,
	}

	if err := services.CreateProject(&project); err != nil {
		c.JSON(500, gin.H{"message" : "failed to create project"})
		return
	}

	resp, _ := services.FindByIDProject(project.ID)

	c.JSON(201, gin.H{"data" : resp})
}

func IndexProject(c *gin.Context){
	project, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"message" : "failed to fetch project"})
	}

	c.JSON(200, gin.H{"data" : project})
}

func ShowProject(c *gin.Context){
	projectID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message" : "invalid Id"})
		return
	}

	project, err := services.FindByIDProject(projectID)
	if err != nil {
		c.JSON(404, gin.H{"message" : "project not found"})
		return
	}

	c.JSON(200, gin.H{"data" : project})
}

func UpdateProject(c *gin.Context){
	projectID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message" : "invalid Id"})
	}

	var req requests.ProjectRequest
	if err := c.ShouldBindJSON(&req); err !=nil {
		c.JSON(400, gin.H{"message" : err.Error()})
	}

	data := map[string]interface{}{}

	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Description != "" {
		data["description"] = req.Description
	}
	if req.WebUrl != "" {
		data["web_url"] = req.WebUrl
	}

	if len(data) == 0 {
		c.JSON(400, gin.H{"message" : "no data to update"})
		return
	}

	if err := services.UpdateProject(projectID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	project, _ := services.FindByIDProject(projectID)
	c.JSON(200, gin.H{"data": project})
}