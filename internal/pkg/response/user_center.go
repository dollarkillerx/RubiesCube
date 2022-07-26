package response

import "github.com/dollarkillerx/RubiesCube/internal/pkg/models"

type UserLogin struct {
	JWT     string         `json:"jwt"`
	Payload models.JSONMap `json:"payload"`
}
