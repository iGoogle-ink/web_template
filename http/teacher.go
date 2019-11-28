package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"web_template/ecode"
)

func teacherAdd(c echo.Context) error {
	//commonSvc.TeacherAdd()
	return JSON(c, nil, nil)
}

func teacherList(c echo.Context) error {
	rsp, err := commonSvc.StudentList()
	return JSON(c, rsp, err)
}

func teacherById(c echo.Context) error {
	pId := c.QueryParam("id")
	if pId == "" {
		return ecode.RequestErr
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		return ecode.RequestErr
	}
	rsp, err := commonSvc.StudentById(id)
	return JSON(c, rsp, err)
}
