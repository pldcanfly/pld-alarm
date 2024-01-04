package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/handler"
	"github.com/pldcanfly/pld-alarm/services"
)

func (s *Server) setRoutes() {
	m := s.media
	s.router.GET("/", MediaHandler(m, handler.HandleRoot))
	s.router.GET("/clock", MediaHandler(m, handler.HandleClock))
	s.router.GET("/feed", MediaHandler(m, handler.HandleNewsFeed))
	s.router.GET("/media", MediaHandler(m, handler.HandleMedia))
}

type MediaHandlerFunc func(c echo.Context, mediaState *services.MediaState) error

func MediaHandler(mediaState *services.MediaState, h MediaHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c, mediaState)
	}
}
