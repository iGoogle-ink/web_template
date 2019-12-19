package hs

import (
	"fmt"
	"testing"
)

func TestDao_StudentList(t *testing.T) {
	stus, err := d.StudentList()
	if err != nil {
		fmt.Println("StudentList err:", err)
		return
	}
	for _, v := range stus {
		fmt.Println("Student:", *v)
	}
}

func TestDao_StudentById(t *testing.T) {
	var (
		id = []int{1, 2, 3, 4}
	)
	stu, err := d.StudentByIds(id)
	if err != nil {
		fmt.Println("StudentByIds err:", err)
		return
	}
	for _, s := range stu {
		fmt.Println("StudentByIds:", s)
	}
}
