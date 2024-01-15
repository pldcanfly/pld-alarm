package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pldcanfly/pld-alarm/services"
)

type Server struct {
	listenAddr string
	router     *echo.Echo
	MediaState *services.MediaState
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		router:     echo.New(),
		MediaState: services.NewMediaState(),
	}
}

func (s *Server) Run() {
	e := s.router
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
	}))
	s.setRoutes()
	e.Start(s.listenAddr)
}
