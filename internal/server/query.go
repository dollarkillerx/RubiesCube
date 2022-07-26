package server

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/errs"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/request"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/response"
	"github.com/dollarkillerx/RubiesCube/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"log"

	"strings"
)

func (s *Server) Query(ctx *gin.Context) {
	model := utils.GetAuthModel(ctx)

	var payload request.Query
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	payload.Search = strings.TrimSpace(payload.Search)
	if payload.Limit >= 20 {
		payload.Limit = 20
	}

	sql := s.storage.DB().
		Model(&models.KVStorage{}).
		Where("project_id = ?", model.ProjectID).
		Where("account = ?", model.Account).
		Where("bucket = ?", payload.Bucket)

	if payload.Key != "" {
		sql = sql.Where("key = ?", payload.Key)
	}

	if payload.Search != "" {
		sql = sql.Where("search like ?", strings.ReplaceAll(`%P%`, "P", payload.Search))
	}

	if payload.Limit != 0 && payload.Offset != 0 {
		sql = sql.Limit(payload.Limit).Offset(payload.Offset)
	}

	var jsonQuery *datatypes.JSONQueryExpression

	if len(payload.Equals) != 0 {
		jsonQuery = datatypes.JSONQuery("payload")
		for _, v := range payload.Equals {
			jsonQuery = jsonQuery.Equals(v.Value, v.Key)
		}
	}

	var kv []models.KVStorage
	if jsonQuery == nil {
		err = sql.Find(&kv).Error
	} else {
		err = sql.Find(&kv, jsonQuery).Error
	}

	if err != nil {
		log.Println(err)
		response.Return(ctx, errs.SqlSystemError)
		return
	}

	response.Return(ctx, kv)
}
