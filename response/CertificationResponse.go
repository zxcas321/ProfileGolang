package response

type CertificationResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Issuer     string `json:"issuer"`
	IssueDate  string `json:"issue_date"`
	ExpiryDate string `json:"expiry_date,omitempty"`
	ImageUrl   string `json:"image_url"`
}
