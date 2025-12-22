package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	projectService "github.com/zxcas321/ProfileGolang/services/project"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateProject(c *gin.Context){
	userID, err := utils.DecodeID(c.Param("userId"))
	if err != nil {
        c.JSON(404, gin.H{"message": "invalid user Id"})
        return
    }

	var req requests.ProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message" : err.Error()})
	}

	project := models.Project{
		UserID : userID,
		Name: req.Name,
		Description: req.Description,
		WebUrl: req.WebUrl,
	}

	if err := projectService.CreateProject(&project); err != nil {
		c.JSON(500, gin.H{"message" : "failed to create project"})
		return
	}

	resp, _ := projectService.FindByIDProject(project.ID)

	c.JSON(201, gin.H{"data" : resp})
}

func IndexProject(c *gin.Context){
	project, err := projectService.FindAllProject()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch project"})
		return
	}

	c.JSON(200, gin.H{"data" : project})
}

func ShowProject(c *gin.Context){
	projectID, err := utils.DecodeID(c.Param("projectId"))
	if err != nil {
		c.JSON(404, gin.H{"message" : "invalid project Id"})
		return
	}

	project, err := projectService.FindByIDProject(projectID)
	if err != nil {
		c.JSON(404, gin.H{"message" : "project not found"})
		return
	}

	c.JSON(200, gin.H{"data" : project})
}

func UpdateProject(c *gin.Context){
	projectID, err := utils.DecodeID(c.Param("projectId"))
	if err != nil {
		c.JSON(404, gin.H{"message" : "invalid project Id"})
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

	if err := projectService.UpdateProject(projectID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	project, _ := projectService.FindByIDProject(projectID)
	c.JSON(200, gin.H{"data": project})
}

func DeleteProject(c *gin.Context) {
	projectID, err := utils.DecodeID(c.Param("projectId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	if err := projectService.DeleteProject(projectID); err != nil {
		c.JSON(404, gin.H{"message": "project not found"})
		return
	}

	c.JSON(200, gin.H{"message" : "Deleted successfully"})
}