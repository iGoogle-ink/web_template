/*
	high school
*/
package http

import (
	"fmt"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"web_template/ecode"
	"web_template/model/hs"

	"github.com/labstack/echo/v4"
)

func studentAdd(c echo.Context) error {
	req := new(hs.StudentAddReq)
	if err := c.Bind(req); err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	return JSON(c, nil, schoolSrv.StudentAdd(req))
}

//var count int32 =0

func studentList(c echo.Context) error {
	//hystrix.Do(config.ProjectName, func() error {
	//	rsp, err := schoolSrv.StudentList()
	//	return JSON(c, rsp, err)
	//}, nil)
	//var (
	//rsp *hs.StudentListRsp
	//err error
	//)
	//atomic.AddInt32(&count,1)
	//fmt.Println(count)
	hystrix.Do(config.ProjectName, func() error {
		rsp, err := schoolSrv.StudentList()
		if err == nil {
			JSON(c, Pager{PageNo: 1, PageSize: 20}.Apply(12, rsp), nil)
		}
		return err
	}, func(e error) error {
		switch e {
		case hystrix.ErrCircuitOpen:
			fmt.Println("studentList 熔断开启:", e)
			JSON(c, nil, ecode.New(502, "熔断开启"))
		case hystrix.ErrMaxConcurrency:
			fmt.Println("studentList 熔断超过最大并发:", e)
			JSON(c, nil, ecode.New(502, "熔断超过最大并发"))
		case hystrix.ErrTimeout:
			fmt.Println("studentList 熔断超时:", e)
			JSON(c, nil, ecode.New(502, "熔断超时"))
		default:
			fmt.Println("studentList default:", e)
			JSON(c, nil, e)
		}
		return e
	})
	return nil
}

func studentById(c echo.Context) error {
	pId := c.QueryParam("id")
	if pId == "" {
		return JSON(c, nil, ecode.RequestErr)
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}

	//hystrix.Do(config.ProjectName, func() error {
	//	rsp, err := schoolSrv.StudentById(id)
	//	JSON(c, rsp, err)
	//	return err
	//}, nil)
	hystrix.Do(config.ProjectName, func() error {
		rsp, err := schoolSrv.StudentById(id)
		if err == nil {
			JSON(c, rsp, nil)
		}
		return err
	}, func(e error) error {
		switch e {
		case hystrix.ErrCircuitOpen:
			fmt.Println("studentList 熔断开启:", e)
			JSON(c, nil, ecode.New(502, "熔断开启"))
		case hystrix.ErrMaxConcurrency:
			fmt.Println("studentList 熔断超过最大并发:", e)
			JSON(c, nil, ecode.New(502, "熔断超过最大并发"))
		case hystrix.ErrTimeout:
			fmt.Println("studentList 熔断超时:", e)
			JSON(c, nil, ecode.New(502, "熔断超时"))
		default:
			JSON(c, nil, e)
		}
		return e
	})
	return nil
}

func teacherAdd(c echo.Context) error {
	req := new(hs.TeacherAddReq)
	if err := c.Bind(req); err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	return JSON(c, nil, schoolSrv.TeacherAdd(req))
}

func teacherList(c echo.Context) error {
	rsp, err := schoolSrv.TeacherList()
	return JSON(c, rsp, err)
}

func teacherById(c echo.Context) error {
	pId := c.QueryParam("id")
	if pId == "" {
		return JSON(c, nil, ecode.RequestErr)
	}
	id, err := strconv.Atoi(pId)
	if err != nil {
		return JSON(c, nil, ecode.RequestErr)
	}
	rsp, err := schoolSrv.TeacherById(id)
	return JSON(c, rsp, err)
}

func sasas() {

}
