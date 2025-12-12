package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	"github.com/zxcas321/ProfileGolang/services"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateCertification(c *gin.Context) {
	var req requests.CertificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	certification := models.Certification{
		Title: req.Title,
		Issuer: req.Issuer,
		IssueDate: req.IssueDate,
		ExpiryDate: req.ExpiryDate,
		ImageUrl: req.ImageUrl,
	}

	if err := services.CreateCertification(&certification); err != nil {
		c.JSON(500, gin.H{"message": "failed to create certification"})
		return
	}
	resp, _ := services.FindByIDCertification(certification.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexCertification(c *gin.Context) {
	certification, err := services.FindAllCertification()
	if err != nil {
		c.JSON(500, gin.H{"message": "failerd to fetch certification"})
		return
	}

	c.JSON(200, gin.H{"data": certification})
}

func ShowCertification(c *gin.Context) {
	certificationID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	certification, err := services.FindByIDCertification(certificationID)
	if err != nil {
		c.JSON(404, gin.H{"message": "certification not found"})
		return
	}

	c.JSON(200, gin.H{"Data": certification})
}

func UpdateCertification(c *gin.Context) {
	certificationID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"Message": "invalid Id"})
	}

	var req requests.CertificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	data := map[string]interface{}{}

	data["title"] = req.Title
	data["issuer"] = req.Issuer
	data["issue_date"] = req.IssueDate
	
	if req.ExpiryDate == nil {
		data["expiry_date"] = nil
	}else{
		data["expiry_date"] = req.ExpiryDate
	}

	if err := services.UpdateExperience(certificationID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

}

func DeleteExperience(c *gin.Context) {
	experienceID, err := utils.DecodeID(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid id"})
		return
	}

	if err := services.DeleteExperience(experienceID); err != nil {
		c.JSON(4040, gin.H{"Message": "experience not found"})
	}

	c.JSON(200, gin.H{"message": "experience deleted successfully"})
}
