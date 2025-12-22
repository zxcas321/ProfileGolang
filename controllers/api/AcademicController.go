package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/requests"
	academicService "github.com/zxcas321/ProfileGolang/services/academic"
	"github.com/zxcas321/ProfileGolang/utils"
)

func CreateAcademic(c *gin.Context) {
	userID, err := utils.DecodeID(c.Param("userId"))
	if err != nil {
        c.JSON(404, gin.H{"message": "invalid user Id"})
        return
    }

	var req requests.AcademicRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	academic := models.Academics{
		UserID : userID,
		Institution : req.Institution,
		Degree : req.Degree,
		Field : req.Field,
		StartDate : req.StartDate,
		EndDate : req.EndDate,
		GPA :req.GPA,
		Description : req.Description,
	}

	if err := academicService.CreateAcademic(&academic); err != nil {
		c.JSON(500, gin.H{"message": "failed to create Academics"})
		return
	}
	resp, _ := academicService.FindByIDAcademic(academic.ID)

	c.JSON(201, gin.H{"data": resp})
}

func IndexAcademic(c *gin.Context) {
	academic, err := academicService.FindAllAcademic()
	if err != nil {
		c.JSON(500, gin.H{"message": "failed to fetch academic"})
		return
	}

	c.JSON(200, gin.H{"data": academic})
}

func ShowAcademic(c *gin.Context) {
	academicID, err := utils.DecodeID(c.Param("academicId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid Id"})
		return
	}

	academic, err := academicService.FindByIDAcademic(academicID)
	if err != nil {
		c.JSON(404, gin.H{"message": "academic not found"})
		return
	}

	c.JSON(200, gin.H{"Data": academic})
}

func UpdateAcademic(c *gin.Context) {
	academicID, err := utils.DecodeID(c.Param("academicId"))
	if err != nil {
		c.JSON(404, gin.H{"Message": "invalid Id"})
	}

	var req requests.AcademicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
	}

	data := map[string]interface{}{}

	data["institution"] = req.Institution
	data["degree"] = req.Degree
	data["Field"] = req.Field
	data["start_date"] = req.StartDate
	data["end_date"] = req.EndDate
	data["GPA"] = req.GPA
	data["description"] = req.Description


	if err := academicService.UpdateAcademic(academicID, data); err != nil {
		c.JSON(500, gin.H{"message": "update failed"})
		return
	}

}

func DeleteAcademic(c *gin.Context) {
	academicID, err := utils.DecodeID(c.Param("academicId"))
	if err != nil {
		c.JSON(404, gin.H{"message": "invalid id"})
		return
	}

	if err := academicService.DeleteAcademic(academicID); err != nil {
		c.JSON(4040, gin.H{"Message": "academic not found"})
	}

	c.JSON(200, gin.H{"message": "academic deleted successfully"})
}
