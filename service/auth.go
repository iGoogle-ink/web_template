package service

import (
	"fmt"
	"net/http"
	"time"

	"web_template/ecode"
	"web_template/model"

	"github.com/labstack/echo/v4"
)

func (s *Service) KeyAuthValidator(session string, c echo.Context) (bool, error) {
	cookie, err := c.Cookie("cookie")
	if err != nil {
		return false, json(c, ecode.CookieErr)
	}
	fmt.Println("cookie:", cookie.String())

	id, err := s.dao.Redis.Get(model.AuthKey + session).Int64()
	if err != nil {
		return false, json(c, ecode.Unauthorized)
	}
	if id == 0 {
		return false, json(c, ecode.Unauthorized)
	}
	s.dao.Redis.Expire(model.AuthKey+session, time.Hour*24)
	return true, nil
}

func json(c echo.Context, err error) error {
	codes := ecode.AnalyseError(err)
	d := &model.ReturnData{
		Code:    codes.Code(),
		Message: codes.Message(),
		Data:    nil,
	}
	return c.JSON(http.StatusOK, d)
}
