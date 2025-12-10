package requests

import "time"

type ExperienceRequest struct {
	CompanyName string    `json:"company_name"`
	Role        string    `json:"role"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
