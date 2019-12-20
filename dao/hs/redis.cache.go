package hs

import (
	"log"
	"strconv"
	"time"

	"web_template/model/hs"
	"web_template/pkg"
)

func (d *Dao) AddCacheTeacherIds(ids []int64) (err error) {
	idStr := pkg.JoinInts(ids)
	if err = d.Redis.Set(_RedisKeyTeacherIdAll, idStr, time.Duration(d.c.ReloadTime+2)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func (d *Dao) AddCacheTeacher(tch *hs.Teacher) (err error) {
	var key = redisTeacherIdKey(int64(tch.Id))
	tchMap := tch.FormatToMap()
	if err = d.Redis.HMSet(key, tchMap).Err(); err != nil {
		return err
	}
	return nil
}

func (d *Dao) AddCacheTeachers(tchs []*hs.Teacher) (err error) {
	// 存 Teacher
	for _, tch := range tchs {
		if inerr := d.AddCacheTeacher(tch); inerr != nil {
			log.Printf("d.AddCacheTeacher(%v) Error(%v)\n", tch, inerr)
			continue
		}
	}
	return nil
}

func (d *Dao) CacheTeacherIds() (ids []int64, err error) {
	idStr, err := d.Redis.Get(_RedisKeyTeacherIdAll).Result()
	if err != nil {
		return nil, err
	}
	ids, err = pkg.SplitInts(idStr)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (d *Dao) CacheTeachers(ids []int64) (tchs []*hs.Teacher, err error) {
	for _, id := range ids {
		tch, err := d.CacheTeacherById(id)
		if err != nil {
			return nil, err
		}
		tchs = append(tchs, tch)
	}
	return tchs, nil
}

func (d *Dao) CacheTeacherById(id int64) (tch *hs.Teacher, err error) {
	var key = redisTeacherIdKey(id)
	mp, err := d.Redis.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	tch = new(hs.Teacher)
	tId, _ := strconv.Atoi(mp["id"])
	tch.Id = tId
	tch.Name = mp["name"]
	tAge, _ := strconv.Atoi(mp["age"])
	tch.Age = tAge
	tGender, _ := strconv.Atoi(mp["gender"])
	tch.Gender = tGender
	tch.Subject = mp["subject"]
	cSec, _ := strconv.ParseInt(mp["ctime"], 10, 64)
	tch.Ctime = time.Unix(cSec, 0)
	mSec, _ := strconv.ParseInt(mp["mtime"], 10, 64)
	tch.Mtime = time.Unix(mSec, 0)
	//ctime, _ := time.Parse("2006-01-02T15:04:05+08:00", mp["ctime"])
	//tch.Ctime = ctime
	//mtime, _ := time.Parse("2006-01-02T15:04:05+08:00", mp["mtime"])
	//tch.Mtime = mtime
	return tch, nil
}
