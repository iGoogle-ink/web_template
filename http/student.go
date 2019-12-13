package http

import (
	"strconv"

	"web_template/ecode"
	"web_template/model/school"

	"github.com/labstack/echo/v4"
)

func studentAdd(c echo.Context) error {
	req := new(school.StudentAddReq)
	if err := c.Bind(req); err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	return JSON(c, nil, schoolSrv.StudentAdd(req))
}

func studentList(c echo.Context) error {
	rsp, err := schoolSrv.StudentList()
	return JSON(c, rsp, err)
}

func studentById(c echo.Context) error {
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
