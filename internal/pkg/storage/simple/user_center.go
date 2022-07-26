package simple

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/pkg/errors"
)

func (s *Simple) GetUserCenter(projectID string, account string) (*models.UserCenter, error) {
	var uc models.UserCenter
	err := s.DB().Model(&models.UserCenter{}).
		Where("account = ?", account).
		Where("project_id = ?", projectID).First(&uc).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &uc, nil
}
