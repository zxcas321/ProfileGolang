package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateAdmin(c *gin.Context) {
	var req requests.AdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	admin := models.Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := services.CreateAdmin(&admin); err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	resp, _ := services.FindByIDAdmin(admin.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexAdmin(c *gin.Context) {
	admin, err := services.FindAllAdmin()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch admin"})
		return
	}

	c.JSON(200, gin.H{"data": admin})
}

func ShowAdmin(c *gin.Context) {
	adminID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	admin, err := services.FindByIDAdmin(adminID)
	if err != nil {
		c.JSON(404, gin.H{"message": "admin not found"})
		return
	}

	c.JSON(200, gin.H{"Data": admin})
}

func UpdateAdmin(c *gin.Context) {
	adminID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"Message": "invalid Id"})
	}

	var req requests.AdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	data := map[string]interface{}{}

	if req.Username != "" {
		data["username"] = req.Username
	}
	if req.Email != "" {
		data["email"] = req.Email
	}

	if err := services.UpdateAdmin(adminID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

}

func DeleteAdmin(c *gin.Context) {
	adminID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid id"})
		return
	}

	if err := services.DeleteAdmin(adminID); err != nil {
		c.JSON(4040, gin.H{"Message": "admin not found"})
	}

	c.JSON(200, gin.H{"message": "admin deleted successfully"})
}
