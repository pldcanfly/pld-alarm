package server

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) setRoutes() {
	s.router.GET("/", ContextHandler(s, HandleRoot))
	s.router.GET("/clock", ContextHandler(s, HandleClock))
	s.router.GET("/feed", ContextHandler(s, HandleNewsFeed))
	s.router.GET("/media", ContextHandler(s, HandleMedia))
	s.router.GET("/media/stop", ContextHandler(s, HandleMediaStop))
	s.router.GET("/media/play", ContextHandler(s, HandleMediaPlay))

	s.router.GET("/alarm", ContextHandler(s, HandleAlarmEvents))
	s.router.GET("/snooze", ContextHandler(s, HandleSnooze))

}

type ContextHandlerFunc func(c echo.Context, s *Server) error

func ContextHandler(s *Server, h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("server", s)
		return h(c)
	}
}
