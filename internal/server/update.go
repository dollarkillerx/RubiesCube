package server

import (
	"github.com/dollarkillerx/RubiesCube/internal/pkg/errs"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/request"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/response"
	"github.com/dollarkillerx/RubiesCube/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) Update(ctx *gin.Context) {
	model := utils.GetAuthModel(ctx)

	var payload request.Update
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	err = s.storage.KvUpdate(model.ProjectID, model.Account, payload.Bucket, payload.Key, payload.Payload)
	if err != nil {
		response.Return(ctx, errs.SqlSystemError)
		return
	}

	response.Return(ctx, gin.H{})
}
