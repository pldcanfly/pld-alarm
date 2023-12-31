package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
)

func getMedia() templ.Component {
	return components.Media()
}

func HandleMedia(c echo.Context) error {
	return Render(c, http.StatusOK, getMedia())
}
