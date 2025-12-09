package requests

import "time"

type certificationRequest struct {
	Title      string     `json:"title"`
	Issuer     string     `json:"issuer"`
	IssueDate  time.Time  `json:"issue_date"`
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	ImageUrl   string     `json:"image_url"`
}
