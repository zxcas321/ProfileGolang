package response

type ExperienceResponse struct {
	ID          string `json:"id"`
	CompanyName string `json:"company_name"`
	Role        string `json:"role"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
