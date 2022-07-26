package request

type AuthJWT struct {
	ProjectID string `json:"project_id"`
	Account   string `json:"account"`
}
