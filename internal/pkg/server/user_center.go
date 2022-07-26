package server

import (
	"fmt"
	"time"

	"github.com/dollarkillerx/RubiesCube/internal/conf"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/errs"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/request"
	"github.com/dollarkillerx/RubiesCube/internal/pkg/response"
	"github.com/dollarkillerx/RubiesCube/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) userCenter(ctx *gin.Context) {
	var payload request.UserLogin
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response.Return(ctx, errs.BadRequest)
		return
	}

	uc, err := s.storage.GetUserCenter(payload.ProjectID, payload.Account)
	if err != nil {
		response.Return(ctx, errs.LoginFailed)
		return
	}

	if uc.Password != utils.Md5Encode(fmt.Sprintf("%s_%s_%s", conf.CONF.Salt, payload.Password, payload.ProjectID)) {
		response.Return(ctx, errs.LoginFailed)
		return
	}

	token, err := utils.JWT.CreateToken(request.AuthJWT{
		ProjectID: uc.ProjectId,
		Account:   uc.Account,
	}, time.Now().Add(time.Hour*24*600).Unix())
	if err != nil {
		response.Return(ctx, errs.SystemError)
		return
	}

	response.Return(ctx, response.UserLogin{
		JWT:     token,
		Payload: uc.Payload,
	})
}
