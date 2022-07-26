package server

import (
	"github.com/dollarkillerx/RubiesCube/internal/conf"
	middleware2 "github.com/dollarkillerx/RubiesCube/internal/middleware"
	"github.com/dollarkillerx/RubiesCube/internal/storage"
	"github.com/dollarkillerx/RubiesCube/internal/storage/simple"
	"github.com/gin-gonic/gin"
)

type Server struct {
	app     *gin.Engine
	storage storage.Interface
}

func NewServer() *Server {
	return &Server{
		app: gin.New(),
	}
}

func (s *Server) Run() error {
	s.router()
	newSimple, err := simple.NewSimple(&conf.CONF.PgSQLConfig)
	if err != nil {
		return err
	}

	s.storage = newSimple

	s.app.Use(middleware2.Cors())
	s.app.Use(middleware2.HttpRecover())
	s.app.Use(middleware2.SetBasicInformation())

	return s.app.Run(conf.CONF.ListenAddr)
}
