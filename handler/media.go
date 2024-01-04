package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
	"github.com/pldcanfly/pld-alarm/services"
)

func getMedia() templ.Component {
	return components.Media(components.MediaAudio)
}

func HandleMedia(c echo.Context, mediaState *services.MediaState) error {
	c.Response().Header().Add("HX-Push-Url", "/media")
	return Render(c, http.StatusOK, Layout(getFeed()))
}
