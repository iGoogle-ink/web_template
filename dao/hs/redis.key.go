package hs

import "fmt"

const (
	_RedisKeyTeacherIdAll = "CACHE_TEACHER_ID_ALL"
	_RedisKeyTeacherId    = "CACHE_TEACHER_%d"
	_RedisKeyTeacher      = "CACHE_TEACHER"
)

func redisTeacherIdKey(id int64) (key string) {
	return fmt.Sprintf(_RedisKeyTeacherId, id)
}
