package response

type SkillsResponse struct{
	ID string `json:"id"`
	SkillName string `json:"skill_name"`
	Category  string `json:"category"`
}