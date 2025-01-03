package utils

import "time"

var BangkokTime *time.Location

func init() {
	BootTimeLocation()
}

func BootTimeLocation() {
	loc, err := time.LoadLocation("Asia/Bangkok")

	BangkokTime = loc

	if err != nil {
		panic(err)
	}
}

func TimeInBangkok(t time.Time) time.Time {
	return t.In(BangkokTime)
}

func TimeNow() time.Time {
	return time.Now().In(BangkokTime)
}

func TimeToday() time.Time {
	now := TimeNow()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, BangkokTime)
}

func TimeNowPtr() *time.Time {
	now := TimeNow()
	return &now
}
