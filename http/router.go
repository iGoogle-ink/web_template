package http

import "github.com/labstack/echo/v4"

func router(e *echo.Echo) {
	e.GET("/ping", ping)
	stu := e.Group("/student" /*keyAuth()*/)
	{
		stu.GET("/list", studentList)
		stu.GET("/id", studentById)
	}
}

func ping(c echo.Context) error {
	return JSON(c, "PONG", nil)
}
