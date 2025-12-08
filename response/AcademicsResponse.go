package response

type AcademicsResponse struct {
	ID          string `json:"id"`
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Field       string `json:"field"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	GPA         string `json:"GPA"`
	Description string `json:"description"`
}
