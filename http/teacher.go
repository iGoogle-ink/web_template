package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"web_template/ecode"
	"web_template/model"
)

func teacherAdd(c echo.Context) error {
	req := new(model.TeacherAddReq)
	if err := c.Bind(req); err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	return JSON(c, nil, commonSvc.TeacherAdd(req))
}

func teacherList(c echo.Context) error {
	rsp, err := commonSvc.TeacherList()
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
