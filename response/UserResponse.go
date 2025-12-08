package response

type UserResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	LinkedinUrl    string `json:"linkedin_url"`
	GithubUrl      string `json:"github_url"`
	Location       string `json:"location"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`

	Experience    []ExperienceResponse    `json:"experience"`
	Project       []ProjectResponse       `json:"project"`
	Academics     []AcademicsResponse     `json:"academics"`
	Certification []CertificationResponse `json:"certification"`
}
