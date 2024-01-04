package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"github.com/pldcanfly/pld-alarm/components"
)

func getFeed() templ.Component {

	u := "https://www.derstandard.at/rss"
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(u)
	if err != nil {
		fmt.Println(err)
	}
	return components.Feed(feed)
}

func HandleNewsFeed(c echo.Context) error {
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	c.Response().Header().Add("HX-Push-Url", "/feed")
	return Render(c, http.StatusOK, Layout(getFeed(), components.StateFeed, s.MediaState))
}
