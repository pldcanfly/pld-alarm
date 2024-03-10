package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleAlarmEvents(c echo.Context) error {

	fmt.Println("/alarms triggered")
	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}

	c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
	c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
	c.Response().Header().Set(echo.HeaderConnection, "keep-alive")
	c.Response().Writer.WriteHeader(http.StatusOK)

	s.AlarmEventContext = c.Response()
	i := 0
	for {
		i++
		fmt.Println(i)
		fmt.Fprintf(c.Response(), "event: test\n")
		fmt.Fprintf(c.Response(), "data: %s %d\n\n", "TESTING THE TESasdfT", i)

		c.Response().Flush()

		time.Sleep(5 * time.Second)
	}

}

func HandleSnooze(c echo.Context) error {

	s, ok := c.Get("server").(*Server)
	if !ok {
		return fmt.Errorf("server context error")
	}
	s.AlarmService.Snooze()

	return nil

}
