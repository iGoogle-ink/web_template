package http

import (
	"strconv"

	"web_template/ecode"
	"web_template/model/hs"

	"github.com/labstack/echo/v4"
)

func teacherAdd(c echo.Context) error {
	req := new(hs.TeacherAddReq)
	if err := c.Bind(req); err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	return JSON(c, nil, schoolSrv.TeacherAdd(req))
}

func teacherList(c echo.Context) error {
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	startInt64, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	endInt64, err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	rsp, err := schoolSrv.TeacherList(startInt64, endInt64)
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
	rsp, err := schoolSrv.StudentById(id)
	return JSON(c, rsp, err)
}
