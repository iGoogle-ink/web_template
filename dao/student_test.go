package dao

import (
	"fmt"
	"testing"
)

func TestDao_StudentList(t *testing.T) {

	stus, err := dao.StudentList()
	if err != nil {
		fmt.Println("StudentList err:", err)
		return
	}
	fmt.Println("StudentList:", stus)
}

func TestDao_StudentById(t *testing.T) {
	var (
		id = 1
	)
	stu, err := dao.StudentById(id)
	if err != nil {
		fmt.Println("StudentById err:", err)
		return
	}
	fmt.Println("StudentById:", stu)
}
