package pkg

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// Time be used to MySql timestamp converting.
type Time int64

// Scan scan time.
func (t *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case time.Time:
		*t = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		fmt.Println("Scan Err:", err)
		*t = Time(i)
	}
	return
}

// Value get time value.
func (t Time) Value() (driver.Value, error) {
	return time.Unix(int64(t), 0), nil
}

// Time get time.
func (t Time) Time() time.Time {
	return time.Unix(int64(t), 0)
}
