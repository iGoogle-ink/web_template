package auth

import (
	"time"

	"web_template/model"
)

func (d *Dao) GetUserIdBySession(session string) (userId int, err error) {
	id, err := d.Redis.Get(model.AuthKey + session).Int64()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (d *Dao) SetUserIdBySession(session string, userId int) {
	d.Redis.Set(model.AuthKey+session, userId, time.Hour*24)
}

func (d *Dao) ExpireUserIdBySession(session string) {
	d.Redis.Expire(model.AuthKey+session, time.Hour*24)
}
