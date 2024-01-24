package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
	"github.com/pldcanfly/pld-alarm/services"
)

func getMedia(ms *services.MediaState) templ.Component {
	return components.Media(ms)
}

func HandleMedia(c echo.Context) error {
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	ms := s.MediaState

	c.Response().Header().Add("HX-Push-Url", "/media")
	return Render(c, http.StatusOK, Layout(getMedia(ms), components.StateAudio, ms))
}

func HandleMediaPlay(c echo.Context) error {
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	ms := s.MediaState
	err := ms.PlayAudio("test.mp3")
	if err != nil {
		return err
	}
	return HandleMedia(c)
}

func HandleMediaStop(c echo.Context) error {
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	ms := s.MediaState
	ms.StopAudio()

	return HandleMedia(c)
}
