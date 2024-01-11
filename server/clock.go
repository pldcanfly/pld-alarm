package server

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
)

func getClock() templ.Component {
	n := time.Now()
	t := n.Format("15:04")
	d := n.Format("Monday, 2 January ")
	return components.Clock(t, d)
}

func HandleClock(c echo.Context) error {
	// s, ok := c.Get("server").(*Server)
	// if !ok {
	// 	return fmt.Errorf("server context error")
	// }
	// fmt.Println(s.MediaState)
	return Render(c, http.StatusOK, getClock())
}
