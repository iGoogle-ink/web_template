package hs

import (
	"fmt"
	"testing"
)

func TestDao_GetBindings(t *testing.T) {
	list, err := d.GetBindings(1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, b := range list {
		fmt.Println("binding:", b)
	}
}

func TestDao_GetBindingsByPid(t *testing.T) {
	list, err := d.GetBindingsByPid(1, 2)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	for _, b := range list {
		fmt.Println("binding:", b)
	}
}
