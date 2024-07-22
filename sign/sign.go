package sign

import (
	"context"
	"fmt"
	"time"

	"github.com/Garfield247/user_sign_go.git/db"
)

const TIME_LAYOUT_DAY = "2006-01-02"

func buildKey(uId int64, year int, month int) string {
	return fmt.Sprintf("project_name:user:sign:%d:%d:%d", year, month, uId)

}

// Sign 签到
func Sign(uId int64, date string) (err error) {
	redisCli := db.GetRedisInstance()
	dateTime, err := time.Parse(TIME_LAYOUT_DAY, date)
	if err != nil {
		return err
	}
	key := buildKey(uId, dateTime.Year(), int(dateTime.Month()))
	v, err := redisCli.SetBit(context.Background(), key, int64(dateTime.Day()), 1).Result()
	if err != nil {
		return err
	}
	fmt.Printf("v: %v\n", v)
	return nil
}

// GetSignStatusByDate 获取当日签到状态
func GetSignStatusByDate(uId int64, date string) (status int, err error) {
	redisCli := db.GetRedisInstance()
	dateTime, err := time.Parse(TIME_LAYOUT_DAY, date)
	if err != nil {
		return 0, err
	}
	key := buildKey(uId, dateTime.Year(), int(dateTime.Month()))
	v, err := redisCli.GetBit(context.Background(), key, int64(dateTime.Day())).Result()
	if err != nil {
		return 0, err
	}
	fmt.Printf("v: %v\n", v)
	return int(v), nil
}
