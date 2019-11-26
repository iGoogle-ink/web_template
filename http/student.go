package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"web_template/ecode"
	"web_template/model"
)

func studentList(c echo.Context) error {
	// todo 封装 echo 的返回
	data := new(model.ReturnData)
	rsp, err := s.StudentList()
	if err != nil {
		data.Code = ecode.NothingFound
		data.Message = err.Error()
		return c.JSON(http.StatusOK, data)
	}
	data.Code = ecode.OK
	data.Message = "success"
	data.Data = rsp
	return c.JSON(http.StatusOK, data)
}

func studentById(c echo.Context) error {
	pId := c.QueryParam("id")
	id, _ := strconv.Atoi(pId)
	data := new(model.ReturnData)
	rsp, err := s.StudentById(id)
	if err != nil {
		data.Code = ecode.NothingFound
		data.Message = err.Error()
		return c.JSON(http.StatusOK, data)
	}
	data.Code = ecode.OK
	data.Message = "success"
	data.Data = rsp
	return c.JSON(http.StatusOK, data)
}
