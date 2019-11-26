package service

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"web_template/conf"
	"web_template/dao"
	"web_template/ecode"
	"web_template/model"
	"web_template/util"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao

	// some data need init when service new
	student map[int]string
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:       c,
		dao:     dao.New(c),
		student: make(map[int]string),
	}

	// some data need init when service new
	s.loadStudent()
	go s.loadStudentProc()
	return s
}

func (s *Service) KeyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   middleware.DefaultSkipper,
		Validator: s.keyAuthValidator,
		KeyLookup: "header:" + util.AuthKey,
	})
}

func (s *Service) keyAuthValidator(session string, c echo.Context) (bool, error) {
	id, err := s.dao.Redis.Get(util.AuthKey + session).Int64()
	if err != nil {
		return false, c.JSON(http.StatusOK, &model.ReturnData{Code: ecode.Unauthorized, Message: "Session Expired"})
	}
	if id == 0 {
		return false, c.JSON(http.StatusOK, &model.ReturnData{Code: ecode.Unauthorized, Message: "Session Expired"})
	}
	s.dao.Redis.Expire(util.AuthKey+session, time.Hour*24)
	return true, nil
}
