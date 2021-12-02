package jtime

import (
	"fmt"
	"time"
)

/*
* add by yjh 211124
* build
 */

func BuildTimeUtils() *TimeUtils {
	return &TimeUtils{
		FormatDate:     "2006-01-02",
		FormatTime:     "15:04:05",
		FormatDateTime: "2006-01-02 15:04:05",
	}
}

/*
* add by yjh 211124
* build
 */
func BuildTimeUtilsNoFmt() *TimeUtils {
	return &TimeUtils{
		FormatDate:     "20060102",
		FormatTime:     "150405",
		FormatDateTime: "20060102150405",
	}
}

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func NowTime() string {
	return time.Now().Format("15:04:05")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func UnixTime() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func BuildTimeCount() *TimeCount {
	return &TimeCount{
		_startTime: time.Now(),
	}
}

// parse timestr to time
// add by yjh
// 211202
func ParseDateTime(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatDateTimeParse(f), s, time.Local)
}

func ParseTime(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatTimeParse(f), s, time.Local)
}

func ParseDate(f string, s string) (time.Time, error) {

	return time.ParseInLocation(_formatDateParse(f), s, time.Local)
}
