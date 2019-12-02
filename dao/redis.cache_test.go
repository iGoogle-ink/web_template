package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestDao_AddCacheTeacher(t *testing.T) {
	tchs, err := dao.TeacherList()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("tchs:", tchs)
	err = dao.AddCacheTeacher(tchs)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}
}

func TestDao_CacheTeacher(t *testing.T) {
	stime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-11-28 19:00:14", time.Local)
	etime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-12-02 17:35:52", time.Local)

	teacher, err := dao.CacheTeacher(stime, etime, 0, 0)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range teacher {
		fmt.Println("v:", v)
	}
}
