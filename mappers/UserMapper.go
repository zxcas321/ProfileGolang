package mappers

import (
	"github.com/zxcas321/ProfileGolang/models"
	"github.com/zxcas321/ProfileGolang/response"
	"github.com/zxcas321/ProfileGolang/utils"
)

func MapUserToResponse(user models.User) response.UserResponse {
	var project []response.ProjectResponse
	for _, p := range user.Project {
		project = append(project, MapProjectToResponse(p))
	}
	var experience []response.ExperienceResponse
	for _, p := range user.Experience {
		experience = append(experience, MapExperienceToResponse(p))
	}
	var academics []response.AcademicsResponse
	for _, p := range user.Academics {
		academics = append(academics, MapAcademicToResponse(p))
	}
	var certification []response.CertificationResponse
	for _, p := range user.Certification {
		certification = append(certification, MapCertificationToResponse(p))
	}

	return response.UserResponse{
		ID:             utils.EncodeID(user.ID),
		Name:           user.Name,
		Email:          user.Email,
		LinkedinUrl:    user.LinkedinUrl,
		GithubUrl:      user.GithubUrl,
		Location:       user.Location,
		Bio:            user.Bio,
		ProfilePicture: user.ProfilePicture,
		Project:        project,
		Experience:     experience,
		Academics:      academics,
		Certification:  certification,
	}
}
