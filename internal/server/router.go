package server

import "github.com/dollarkillerx/RubiesCube/internal/middleware"

func (s *Server) router() {
	s.app.POST("/api/v1/login", s.userCenter)

	v1 := s.app.Group("/api/v1", middleware.UAAuthorization())
	{
		v1.POST("/storage", s.Storage)
		v1.POST("/query", s.Query)
		v1.POST("/update", s.Update)
		v1.POST("/delete", s.Delete)
	}

	//maV1 := s.app.Group("/api/ma/v1")
	//{
	//	user := maV1.Group("user_center")
	//	{
	//		//user.POST("register")
	//		//user.POST("disabled")
	//		//user.POST("change_password")
	//	}
	//}
}
