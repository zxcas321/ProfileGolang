package requests

type ProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	WebUrl      string `json:"web_url"`
}
