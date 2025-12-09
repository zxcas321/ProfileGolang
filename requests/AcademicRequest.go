package requests

import "time"

type AcademicRequest struct {
	Institution string    `json:"institution"`
	Degree      string    `json:"degree"`
	Field       string    `json:"field"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	GPA         string    `json:"GPA"`
	Description string    `json:"description"`
}
