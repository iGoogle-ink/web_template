package http

import (
	"net/http"

	"web_template/conf"
	"web_template/ecode"
	"web_template/model"
	"web_template/service/auth"
	"web_template/service/pay"
	"web_template/service/school"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"xorm.io/xorm"
)

var (
	schoolSrv *school.Service
	authSrv   *auth.Service
	paySrv    *pay.Service
	config    *conf.Config
)

type Server struct {
	SchoolSrv *school.Service
	AuthSrv   *auth.Service
	PaySrv    *pay.Service
	// some other Service
}

func Init(c *conf.Config, db *xorm.Engine, rds *redis.Client) {
	initService(c, db, rds)
	e := echo.New()
	e.Use(middleware.CORS(), middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip = ${remote_ip}, method = ${method}, uri = ${uri}, Session = ${header:Session}, status = ${status}.\n",
	}))
	router(e)
	if err := e.Start(c.HttpServer.Port); err != nil {
		panic(err)
	}
}

func initService(c *conf.Config, db *xorm.Engine, rds *redis.Client) {
	config = c
	schoolSrv = school.New(c, db, rds)
	authSrv = auth.New(c, rds)
	paySrv = pay.New(c, db, rds)
}

func keyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Skipper:   middleware.DefaultSkipper,
		Validator: authSrv.KeyAuthValidator,
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
