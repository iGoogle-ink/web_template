package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func studentList(c echo.Context) error {
	rsp, err := commonSvc.StudentList()
	return JSON(c, rsp, err)
}

func studentById(c echo.Context) error {
	pId := c.QueryParam("id")
	id, _ := strconv.Atoi(pId)
	rsp, err := commonSvc.StudentById(id)
	return JSON(c, rsp, err)
}
