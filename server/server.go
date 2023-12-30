package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	listenAddr string
	router     *echo.Echo
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		router:     echo.New(),
	}
}

func (s *Server) Run() {
	e := s.router
	// TODO: rethink build so it does not require the static folder
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
	}))
	s.setRoutes()
	e.Start(s.listenAddr)
}
