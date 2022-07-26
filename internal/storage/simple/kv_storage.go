package simple

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/pkg/errors"
	"log"
)

func (s *Simple) StorageKV(kvStorage models.KVStorage) error {
	err := s.DB().Model(&models.KVStorage{}).Create(&kvStorage).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Simple) KvUpdate(project string, account string, bucket string, key string, payload models.JSONMap) (err error) {
	begin := s.DB().Begin()

	defer func() {
		if err == nil {
			e := begin.Commit().Error
			if e != nil {
				log.Println(err)
			}
		} else {
			e := begin.Rollback().Error
			if e != nil {
				log.Println(err)
			}
		}
	}()

	err = begin.
		Model(&models.KVStorage{}).
		Where("project_id = ?", project).
		Where("account = ?", account).
		Where("bucket = ?", bucket).
		Where("key = ?", key).
		Update("payload", payload).Error
	if err != nil {
		return
	}

	return
}

func (s *Simple) KvDelete(project string, account string, bucket string, key string) (err error) {
	begin := s.DB().Begin()

	defer func() {
		if err == nil {
			e := begin.Commit().Error
			if e != nil {
				log.Println(err)
			}
		} else {
			e := begin.Rollback().Error
			if e != nil {
				log.Println(err)
			}
		}
	}()

	err = begin.
		Model(&models.KVStorage{}).
		Where("project_id = ?", project).
		Where("account = ?", account).
		Where("bucket = ?", bucket).
		Where("key = ?", key).Error
	if err != nil {
		return
	}

	return
}
