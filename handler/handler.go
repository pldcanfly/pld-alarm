package handler

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/pldcanfly/pld-alarm/components"
	"github.com/pldcanfly/pld-alarm/services"
)

func Render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}

func Layout(content templ.Component) templ.Component {

	return components.Layout(content, getClock())
}

func getNav(state int) templ.Component {
	return components.Navigation(state)
}

func HandleRoot(c echo.Context, mediaState *services.MediaState) error {
	return Render(c, http.StatusOK, Layout(getFeed()))
}
