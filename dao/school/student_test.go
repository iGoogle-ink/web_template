package school

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
		id = 1
	)
	stu, err := d.StudentById(id)
	if err != nil {
		fmt.Println("StudentById err:", err)
		return
	}
	fmt.Println("StudentById:", *stu)
}
