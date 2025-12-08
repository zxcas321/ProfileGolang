package mappers

import (
	"time"

	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapCertificationToResponse(certification models.Certification) response.CertificationResponse{
	return response.CertificationResponse{
		ID: utils.EncodeID(certification.ID),
		Title: certification.Title,
		Issuer: certification.Issuer,
		IssueDate: certification.IssueDate.Format(time.RFC3339),
		ExpiryDate: certification.ExpiryDate.Format(time.RFC3339),
		ImageUrl: certification.ImageUrl,
	}
}