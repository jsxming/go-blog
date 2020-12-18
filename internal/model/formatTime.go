package model

import (
	"blog/pkg/util"
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// JSONTime format json time field by myself
type JSONTime time.Time

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("JSONTime.UnmarshalJSON Error")
	}
	*t = JSONTime(time.Unix(i/1000, 0))
	return err
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	timestamp := time.Time(t).Unix()
	str := strconv.Itoa(int(timestamp * 1000))
	return []byte(str), nil
}

func (t JSONTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(util.TimeLayout), nil
}

func (t *JSONTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = JSONTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *JSONTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}
