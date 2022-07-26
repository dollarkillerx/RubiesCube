package server

func (s *Server) router() {
	s.app.POST("/search")
}
