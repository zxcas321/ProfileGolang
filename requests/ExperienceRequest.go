package requests

import "time"

type Experience struct {
	CompanyName string    `json:"company_name"`
	Role        string    `json:"role"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
