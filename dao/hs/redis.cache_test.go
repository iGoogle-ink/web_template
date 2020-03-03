package hs

import (
	"fmt"
	"strings"
	"testing"
)

func TestDao_AddCacheTeacher(t *testing.T) {
	tchs, err := d.TeacherList()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("tchs:", tchs)
	err = d.AddCacheTeachers(tchs)
	if err != nil {
		fmt.Println("err2:", err)
		return
	}
}

func TestDao_CacheTeacherIds(t *testing.T) {
	ids, err := d.CacheTeacherIds()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("ids:", ids)
}

func TestDao_CacheTeacherById(t *testing.T) {
	tch, err := d.CacheTeacherById(1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(tch)
}

func TestSss(t *testing.T) {
	var sb strings.Builder
	sb.WriteString(" asdd ")
	temp := strings.TrimSpace(sb.String())
	sb.Reset()
	sb.WriteString(temp)
	fmt.Println("::", sb.String())

}
