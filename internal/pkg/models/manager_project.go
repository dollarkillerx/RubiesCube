package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ManagerProject struct {
	BasicModel
	UserID      string `gorm:"type:varchar(300);index"`
	ProjectName string `gorm:"type:varchar(300);"`
	Payload     ProjectPayload
}

type ProjectPayload struct {
	Buckets []string `json:"buckets"`
}

func (c ProjectPayload) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *ProjectPayload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &c)
}

// GormDBDataType gorm db data type
func (ProjectPayload) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
