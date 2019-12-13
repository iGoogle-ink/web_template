package school

import "fmt"

const (
	_RedisKeyTeacherId = "CACHE_TEACHER_%d"
	_RedisKeyTeacher   = "CACHE_TEACHER"
)

func redisTeacherIdKey(id int) (key string) {
	return fmt.Sprintf(_RedisKeyTeacherId, id)
}
