package server

import "github.com/pldcanfly/pld-alarm/handler"

func (s *Server) setRoutes() {
	s.router.GET("/", handler.HandleRoot)
	s.router.GET("/clock", handler.HandleClock)
	s.router.GET("/feed", handler.HandleNewsFeed)
	s.router.GET("/media", handler.HandleMedia)
}
