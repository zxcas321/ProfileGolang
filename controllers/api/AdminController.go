package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
)

func CreateAdmin(c *gin.Context) {
	var req requests.AdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	admin := models.Admin{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := services.CreateAdmin(&admin); err != nil {
		c.JSON(500, gin.H{
			"message": "failed to create Admin",
		})
		return
	}
	resp, _ := services.FindByIDAdmin(admin.ID)

	c.JSON(201, gin.H{
		"data": resp,
	})
}

func IndexAdmin(c *gin.Context) {
	admin, err := services.FindAllAdmin()
	if err != nil {
		c.JSON(500, gin.H{
			"message" : "failerd to fetch admin",
		})
		return
	}

	c.JSON(200, gin.H{
		"data" : admin,
	})
}
