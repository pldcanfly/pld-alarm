package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
	"github.com/pldcanfly/pld-alarm/services"
	"golang.org/x/net/websocket"
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
	ws := s.MediaWS
	if ws != nil {

		websocket.Message.Send(s.MediaWS, "Handle MediMediaa")
	}
	ms.PlayAudio("test.mp3")
	c.Response().Header().Add("HX-Push-Url", "/media")
	return Render(c, http.StatusOK, Layout(getMedia(ms), components.StateAudio, ms))
}

type PlayerMessage struct {
	Sender string `json:"sender"`
	Action string `json:"action"`
	Data   struct {
		CurrentTime float64 `json:"currentTime`
		Duration    float64 `json:"duration"`
	} `json:"data"`
}

func HandleMediaWS(c echo.Context) error {
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		s.MediaWS = ws

		for {
			var msg PlayerMessage
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {

				fmt.Printf("ws: read %v\n", err)
				return
			}
			fmt.Printf("%s\n", msg)
			s.MediaState.Current = msg.Data.CurrentTime
			s.MediaState.Duration = msg.Data.Duration
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
