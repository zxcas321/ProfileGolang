package requests

type ProjectRequest struct {
	UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WebUrl      string `json:"web_url"`
}
