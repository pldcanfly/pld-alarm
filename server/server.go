package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pldcanfly/pld-alarm/services"
)

type Server struct {
	listenAddr        string
	router            *echo.Echo
	MediaState        *services.MediaState
	AlarmService      *services.Alarm
	AlarmEventContext *echo.Response
}

func NewServer(listenAddr string) *Server {
	ms := services.NewMediaState()
	return &Server{
		listenAddr:   listenAddr,
		router:       echo.New(),
		MediaState:   ms,
		AlarmService: services.NewAlarm(8, 0, "test.mp3", ms),
	}
}

func (s *Server) Run() {
	e := s.router
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
	}))
	s.setRoutes()
	s.AlarmService.SetActive(true)
	e.Start(s.listenAddr)
}
