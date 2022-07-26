package server

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/errs"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/models"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/request"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/response"
	"github.com/dollarkillerx/RubiesCube/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func (s *Server) Storage(ctx *gin.Context) {
	model := utils.GetAuthModel(ctx)

	var payload request.Storage
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	kv := models.KVStorage{
		BasicModel: models.BasicModel{
			ID: xid.New().String(),
		},
		ProjectId: model.ProjectID,
		Account:   model.Account,
		Bucket:    payload.Bucket,
		Key:       payload.Key,
		Search:    payload.Search,
		Payload:   payload.Payload,
	}

	err = s.storage.StorageKV(kv)
	if err != nil {
		response.Return(ctx, errs.SqlSystemError)
		return
	}

	response.Return(ctx, gin.H{})
}
