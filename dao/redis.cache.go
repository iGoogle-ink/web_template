package dao

import (
	"encoding/json"
	"strconv"
	"time"

	"web_template/model"

	"github.com/go-redis/redis/v7"
)

func (d *Dao) CacheTeacherByScore(stime, etime time.Time, offset, count int64) (tchs []*model.Teacher, err error) {
	var results []string
	err = d.Redis.ZRangeByScore(_RedisKeyTeacher, &redis.ZRangeBy{
		Min:    strconv.FormatInt(stime.Unix(), 10),
		Max:    strconv.FormatInt(etime.Unix(), 10),
		Offset: offset,
		Count:  count,
	}).ScanSlice(&results)
	if err != nil {
		return nil, err
	}
	for _, t := range results {
		teacher := new(model.Teacher)
		err := json.Unmarshal([]byte(t), teacher)
		if err != nil {
			return nil, err
		}
		tchs = append(tchs, teacher)
	}
	return tchs, nil
}
func (d *Dao) CacheTeacher(start, end int64) (tchs []*model.Teacher, err error) {
	var results []string
	err = d.Redis.ZRange(_RedisKeyTeacher, start, end).ScanSlice(&results)
	if err != nil {
		return nil, err
	}
	for _, t := range results {
		teacher := new(model.Teacher)
		err := json.Unmarshal([]byte(t), teacher)
		if err != nil {
			return nil, err
		}
		tchs = append(tchs, teacher)
	}
	return tchs, nil
}

func (d *Dao) AddCacheTeacher(tchs []*model.Teacher) (err error) {
	var mems []*redis.Z
	for _, tch := range tchs {
		marshal, err := json.Marshal(tch)
		if err != nil {
			return err
		}
		mems = append(mems, &redis.Z{
			Score:  float64(tch.Mtime.Unix()),
			Member: marshal,
		})
	}
	if err = d.Redis.ZAdd(_RedisKeyTeacher, mems...).Err(); err != nil {
		return err
	}
	if err = d.Redis.Expire(_RedisKeyTeacher, time.Duration(d.c.RedisExpire)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}
