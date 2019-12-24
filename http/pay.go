package http

import "github.com/labstack/echo/v4"

func paySend(c echo.Context) error {
	err := paySrv.SendMessage("我是消息")
	return JSON(c, nil, err)
}
