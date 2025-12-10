package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateUser(c *gin.Context) {
	var req requests.UserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Name:           req.Name,
		Email:          req.Email,
		LinkedinUrl:    req.LinkedinURL,
		GithubUrl:      req.GithubURL,
		Location:       req.Location,
		Bio:            req.Bio,
		ProfilePicture: req.ProfilePicture,
	}

	if err := services.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{
			"message": "failed to create user",
		})
		return
	}

	resp, _ := services.FindByIDUser(user.ID)

	c.JSON(201, gin.H{
		"data": resp,
	})
}

func IndexUser(c *gin.Context) {
	users, err := services.FindAllUser()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to fetch users",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}

func ShowUser(c *gin.Context) {
	decodedID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "invalid Id",
		})
		return
	}

	user, err := services.FindByIDUser(decodedID)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	userID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "invalid Id",
		})
		return
	}

	var req requests.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	data := map[string]interface{}{}

	if req.Name != "" {
		data["name"] = req.Name
	}
	if req.Bio != "" {
		data["bio"] = req.Bio
	}
	if req.Location != "" {
		data["location"] = req.Location
	}

	if len(data) == 0 {
		c.JSON(400, gin.H{"message": "no data to update"})
		return
	}

	if err := services.UpdateUser(userID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

	user, _ := services.FindByIDUser(userID)
	c.JSON(200, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	userID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "invalid Id",
		})
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(404, gin.H{"message": "user not found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "user deleted successfully",
	})
}
