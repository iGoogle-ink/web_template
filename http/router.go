package http

import "github.com/labstack/echo/v4"

func router(e *echo.Echo) {
	e.GET("/ping", ping)
	stu := e.Group("/student" /*, keyAuth()*/) // keyAuth add if you need
	{
		stu.POST("/add", studentAdd)
		stu.GET("/list", studentList)
		stu.GET("/id", studentById)
	}
	tch := e.Group("/teacher")
	{
		tch.POST("/add", teacherAdd)
		tch.GET("/list", teacherList)
		tch.GET("/id", teacherById)
	}
	pay := e.Group("/pay")
	{
		pay.POST("/send", paySend)
	}
}

func ping(c echo.Context) error {
	return JSON(c, "PONG", nil)
}
