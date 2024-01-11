package server

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) setRoutes() {
	s.router.GET("/", ContextHandler(s, HandleRoot))
	s.router.GET("/clock", ContextHandler(s, HandleClock))
	s.router.GET("/feed", ContextHandler(s, HandleNewsFeed))
	s.router.GET("/media", ContextHandler(s, HandleMedia))

	s.router.GET("/ws/media", ContextHandler(s, HandleMediaWS))

}

type ContextHandlerFunc func(c echo.Context, s *Server) error

func ContextHandler(s *Server, h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("server", s)
		return h(c)
	}
}
