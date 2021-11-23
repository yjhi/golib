package jutils

/***************************************************
*    Create By YJH 20210410
*
*
*****************************************************/

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	DateTime string
	Date     string
	Time     string
}

func New() *Time {

	t := &Time{
		DateTime: "2006-01-02 15:04:05",
		Date:     "2006-01-02",
		Time:     "15:04:05",
	}

	return t
}

//兼容老版本
func NewTime() *Time {

	return New()
}

func (t *Time) WithDateTimeFmt(s string) *Time {
	t.SetDateTimeFormat(s)
	return t
}

func (t *Time) WithDateFmt(s string) *Time {
	t.SetDateFormat(s)
	return t
}

func (t *Time) WithTimeFmt(s string) *Time {
	t.SetTimeFormat(s)
	return t
}

func (t *Time) SetDateTimeFormat(s string) {

	s = strings.Replace(s, "yyyy", "2006", 1)
	s = strings.Replace(s, "MM", "01", 1)
	s = strings.Replace(s, "dd", "02", 1)
	s = strings.Replace(s, "HH", "15", 1)
	s = strings.Replace(s, "mm", "04", 1)
	s = strings.Replace(s, "ss", "05", 1)

	t.DateTime = s
}

func (t *Time) SetDateFormat(s string) {

	s = strings.Replace(s, "yyyy", "2006", 1)
	s = strings.Replace(s, "MM", "01", 1)
	s = strings.Replace(s, "dd", "02", 1)

	t.Date = s
}

func (t *Time) SetTimeFormat(s string) {

	s = strings.Replace(s, "HH", "15", 1)
	s = strings.Replace(s, "mm", "04", 1)
	s = strings.Replace(s, "ss", "05", 1)

	t.Time = s
}

func (t *Time) GetDateTime() string {
	return time.Now().Format(t.DateTime)
}

func (t *Time) GetDate() string {
	return time.Now().Format(t.Date)
}

func (t *Time) GetTime() string {
	return time.Now().Format(t.Time)
}

func (*Time) GetUnixTime() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}
