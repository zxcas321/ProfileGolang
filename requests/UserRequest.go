package requests

type UserRequest struct {
	Name           string `json:"name" binding:"required,min=3"`
	Email          string `json:"email" binding:"required,email"`
	LinkedinURL    string `json:"linkedin_url"`
	GithubURL      string `json:"github_url"`
	Location       string `json:"location"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profile_picture"`
}
