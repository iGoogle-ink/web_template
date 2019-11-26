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
	s *service.Service
)

func Init(c *conf.Config) {
	initService(c)
	e := echo.New()
	e.Use(middleware.Recover(), middleware.CORS(), middleware.CSRF(), middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip = ${remote_ip}, method = ${method}, uri = ${uri}, Session = ${header:Session}, status = ${status}.\n",
	}))
	router(e)
	if err := e.Start(c.HTTP.Port); err != nil {
		panic(err)
	}
}

func router(e *echo.Echo) {
	e.GET("/ping", ping)
	stu := e.Group("/student" /*s.KeyAuth()*/)
	{
		stu.GET("/list", studentList)
		stu.GET("/id", studentById)
	}
}

func initService(c *conf.Config) {
	s = service.New(c)
}

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &model.ReturnData{
		Code:    ecode.OK,
		Message: "Ping Ok",
	})
}
