package request

type UserLogin struct {
	ProjectID string `json:"project_id"`
	Account   string `json:"account"`
	Password  string `json:"password"`
}
