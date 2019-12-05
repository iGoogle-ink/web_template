package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"web_template/conf"
	"web_template/ecode"
	"web_template/model"
	"web_template/service"
)

var (
	commonSvc *service.Service
	config    *conf.Config
)

type Server struct {
	Service *service.Service

	// some other Service
}

func (s *Server) Close() {
	s.Service.Close()
}

func Init(c *conf.Config, svr *Server) {
	initService(c, svr)
	e := echo.New()
	e.Use(middleware.CORS(), middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip = ${remote_ip}, method = ${method}, uri = ${uri}, Session = ${header:Session}, status = ${status}.\n",
	}))
	router(e)
	if err := e.Start(c.HTTP.Port); err != nil {
		panic(err)
	}
}

func initService(c *conf.Config, svr *Server) {
	commonSvc = svr.Service
	config = c
}

func keyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   middleware.DefaultSkipper,
		Validator: commonSvc.KeyAuthValidator,
		KeyLookup: "header:" + model.AuthKey,
	})
}

func JSON(c echo.Context, data interface{}, err error) error {
	codes := ecode.AnalyseError(err)
	d := &model.ReturnData{
		Code:    codes.Code(),
		Message: codes.Message(),
		Data:    data,
	}
	return c.JSON(http.StatusOK, d)
}

func String(c echo.Context, statusCode int, msg string) error {
	return c.String(statusCode, msg)
}
