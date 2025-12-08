package response

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WebUrl      string `json:"web_url"`

	Skills []SkillsResponse `json:"skills"`
}
