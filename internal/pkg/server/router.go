package server

func (s *Server) router() {
	s.app.POST("/api/v1/login", s.userCenter)

	v1 := s.app.Group("/api/v1")
	{
		v1.POST("/add")
		v1.POST("/query")
		v1.POST("/update")
		v1.POST("/delete")
	}

	maV1 := s.app.Group("/api/ma/v1")
	{
		user := maV1.Group("user_center")
		{
			user.POST("register")
			user.POST("disabled")
			user.POST("change_password")
		}
	}
}
