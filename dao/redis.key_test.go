package dao

import (
	"testing"
)

func TestDao_redisTeacherIdKey(t *testing.T) {
	key := redisTeacherIdKey(2)
	println(key)
}
