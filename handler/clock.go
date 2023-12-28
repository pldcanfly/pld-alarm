package handler

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
)

func getClock() templ.Component {
	n := time.Now()
	t := n.Format("15:04:05")
	d := n.Format("Monday, 2 January ")
	return components.Clock(t, d)
}

func HandleClock(c echo.Context) error {
	return Render(c, http.StatusOK, getClock())
}
