package jtime

import (
	"fmt"
	"strings"
	"time"
)

type TimeUtils struct {
	FormatDate     string
	FormatTime     string
	FormatDateTime string
}

func _formatDateParse(sf string) string {
	sf = strings.Replace(sf, "yyyy", "2006", 1)
	sf = strings.Replace(sf, "MM", "01", 1)
	sf = strings.Replace(sf, "dd", "02", 1)

	return sf
}

func _formatTimeParse(sf string) string {

	sf = strings.Replace(sf, "hh", "15", 1)
	sf = strings.Replace(sf, "mm", "04", 1)
	sf = strings.Replace(sf, "ss", "05", 1)

	return sf
}

func _formatDateTimeParse(sf string) string {
	sf = _formatDateParse(sf)
	sf = _formatTimeParse(sf)
	return sf
}

func _timeStr(f string) string {
	return time.Now().Format(f)
}

func (t TimeUtils) Time() string {
	return _timeStr(t.FormatTime)
}

func (t TimeUtils) Date() string {
	return _timeStr(t.FormatDate)
}

func (t TimeUtils) DateTime() string {
	return _timeStr(t.FormatDateTime)
}

func (t TimeUtils) TimeWithFmt(f string) string {
	return _timeStr(_formatTimeParse(f))
}
func (t TimeUtils) DateWithFmt(f string) string {
	return _timeStr(_formatDateParse(f))
}
func (t TimeUtils) DateTimeWithFmt(f string) string {
	return _timeStr(_formatDateTimeParse(f))
}

func BuildTimeUtils() *TimeUtils {
	return &TimeUtils{
		FormatDate:     "2006-01-02",
		FormatTime:     "15:04:05",
		FormatDateTime: "2006-01-02 15:04:05",
	}
}

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
