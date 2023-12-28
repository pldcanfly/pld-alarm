package handler

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/mmcdole/gofeed"
	"github.com/pldcanfly/pld-alarm/components"
)

func getFeed() templ.Component {
	return components.Feed()
}

func HandleNewsFeed(c echo.Context) error {
	u := "https://www.derstandard.at/rss"
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", feed)
	return Render(c, http.StatusOK, getFeed())
}
