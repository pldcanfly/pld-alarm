package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
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

		websocket.Message.Send(s.MediaWS, "Handle MediaWebsocket")
	}
	ms.PlayAudio("test.mp3")
	c.Response().Header().Add("HX-Push-Url", "/media")
	return Render(c, http.StatusOK, Layout(getMedia(ms), components.StateAudio, ms))
}

type WSMessage struct {
	Sender string      `json:"sender"`
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}
type MessagePlaying struct {
	CurrentTime float64 `json:"currentTime"`
	Duration    float64 `json:"duration"`
}

type MessageTimeUpdate struct {
	CurrentTime float64 `json:"currentTime"`
	Duration    float64 `json:"duration"`
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
			var msg WSMessage
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				fmt.Printf("mediaws: %v", err)

				return
			}

			switch msg.Action {
			case "playing":
				var msgPlaying MessagePlaying
				data, ok := msg.Data.(map[string]interface{})
				if !ok {
					fmt.Printf("mediaws: playing: invalid message format")
				}

				if err := mapstructure.Decode(data, &msgPlaying); err != nil {

					fmt.Printf("mediaws: playing: invalid data format")
				}

				fmt.Println(msgPlaying)
			case "timeupdate":
				var msgTimeupdate MessageTimeUpdate
				data, ok := msg.Data.(map[string]interface{})
				if !ok {
					fmt.Printf("mediaws: timeupdate: invalid message format")
				}

				if err := mapstructure.Decode(data, &msgTimeupdate); err != nil {

					fmt.Printf("mediaws: timeupdate: invalid data format")
				}

				fmt.Println(msgTimeupdate)

			}

			// websocket.JSON.Send(ws, PlayerMessage{
			// 	Sender: "MediaState",
			// 	Action: "update",
			// })
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
