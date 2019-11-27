package service

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"web_template/ecode"
	"web_template/model"
	"web_template/pkg"
)

func (s *Service) KeyAuthValidator(session string, c echo.Context) (bool, error) {
	id, err := s.dao.Redis.Get(pkg.AuthKey + session).Int64()
	if err != nil {
		return false, json(c, ecode.Unauthorized)
	}
	if id == 0 {
		return false, json(c, ecode.Unauthorized)
	}
	s.dao.Redis.Expire(pkg.AuthKey+session, time.Hour*24)
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
