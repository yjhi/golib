package jtime

import "time"

/*
* add by yjh 211124
* parse time utils
 */
type TimeUtils struct {
	FormatDate     string
	FormatTime     string
	FormatDateTime string
}

/*
* add by yjh 211124
* parse time
 */
func (t TimeUtils) Time() string {
	return _timeStr(t.FormatTime)
}

/*
* add by yjh 211124
* parse date
 */
func (t TimeUtils) Date() string {
	return _timeStr(t.FormatDate)
}

/*
* add by yjh 211124
* parse datetime
 */

func (t TimeUtils) DateTime() string {
	return _timeStr(t.FormatDateTime)
}

/*
* add by yjh 211124
* parse datetime with fmt
 */

func (t TimeUtils) TimeWithFmt(f string) string {
	return _timeStr(_formatTimeParse(f))
}
func (t TimeUtils) DateWithFmt(f string) string {
	return _timeStr(_formatDateParse(f))
}
func (t TimeUtils) DateTimeWithFmt(f string) string {
	return _timeStr(_formatDateTimeParse(f))
}

/*
* add by yjh 211125
* count time
 */
type TimeCount struct {
	_startTime time.Time
	_stopTime  time.Time
}

func (t *TimeCount) Start() {
	t._startTime = time.Now()
}

func (t *TimeCount) Minutes() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Minutes()
}
func (t *TimeCount) Seconds() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Seconds()
}
func (t *TimeCount) Hours() float64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Hours()
}

func (t *TimeCount) Microseconds() int64 {
	t._stopTime = time.Now()

	return t._stopTime.Sub(t._startTime).Microseconds()
}
