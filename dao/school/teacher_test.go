package school

import (
	"fmt"
	"testing"

	"web_template/model/hs"
)

func TestDao_TeacherList(t *testing.T) {
	list, err := d.TeacherList()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, v := range list {
		fmt.Println("teacher:", v)
	}
}

func TestDao_TeacherById(t *testing.T) {
	tch, err := d.TeacherById(2)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("teacher:", tch)
}

func TestDao_TeacherExistById(t *testing.T) {
	has, err := d.TeacherExistById(3)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("teacher_has:", has)
}

func TestDao_TeacherInsert(t *testing.T) {
	tch := &hs.Teacher{
		Name:    "UT测试",
		Subject: "UT测试学科",
	}
	id, err := d.TeacherInsert(tch)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("teacher_id:", id)
}
