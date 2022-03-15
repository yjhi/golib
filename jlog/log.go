package jlog

import (
	"fmt"
	"io"
	"os"

	"gihub.com/yjhi/golib/jtime"
)

/*
* LogUtils
* add by yjh 211123
*
 */

var _timeUtil *jtime.TimeUtils = nil

func _getTimeUtil() *jtime.TimeUtils {
	if _timeUtil == nil {
		_timeUtil = jtime.BuildTimeUtilsNoFmt()
	}

	return _timeUtil
}

func _logDayName(tail string) string {

	logname := _getTimeUtil().Date()
	return logname + "." + tail
}

func _logMonthName(tail string) string {

	logname := _getTimeUtil().DateWithFmt("yyyyMM")
	return logname + "." + tail
}

type LogUtils struct {
	_logFile   *os.File
	_logPath   string
	_logName   string
	LastError  error
	_fullPath  string
	_monthLog  bool
	_dayLog    bool
	_singleLog bool
	_logTail   string
}

func _buildLogUtils(path string, name string, tail string, day bool, month bool, single bool) *LogUtils {

	l := &LogUtils{
		_logPath:   path,
		_logName:   name,
		_logFile:   nil,
		LastError:  nil,
		_fullPath:  "",
		_singleLog: single,
		_dayLog:    day,
		_monthLog:  month,
		_logTail:   tail,
	}

	return l
}

func BuildDayLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, true, false, false)
}
func BuildMonthLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, false, true, false)
}

func BuildSingleLogUtils(path string, name string, tail string) *LogUtils {
	return _buildLogUtils(path, name, tail, false, false, true)
}

/*
* create file function
* 211123
 */
func (l *LogUtils) _openLog(flag string) bool {

	l._logFile, l.LastError = os.OpenFile(l._fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if l.LastError != nil {
		fmt.Println("failed to create file "+flag, l.LastError.Error())
		return false
	}

	return true
}

/*
* write file function
* 211123
 */
func (l *LogUtils) _writeLog(flag string, content string) {
	if _, l.LastError = io.WriteString(l._logFile, jtime.NowDateTime()+" ["+flag+"] "+content+"\n"); l.LastError != nil {
		fmt.Println("failed to write file", l.LastError.Error())
	}
}

func (l *LogUtils) _buildLogFullPath() string {

	var logfile string

	if l._singleLog {
		logfile = l._logPath + "/" + l._logName + "." + l._logTail
	} else if l._monthLog {
		logfile = l._logPath + "/" + l._logName + "_" + _logMonthName(l._logTail)
	} else {
		logfile = l._logPath + "/" + l._logName + "_" + _logDayName(l._logTail)
	}

	return logfile
}

func (l *LogUtils) _appendContent(flag string, content string) {

	logfile := l._buildLogFullPath()

	if l._logFile == nil {

		l._fullPath = logfile

		l._openLog("01")

		l._writeLog(flag, "--------------------First Init LogFile-------------------------")

	} else {

		if l._fullPath == logfile {

		} else {

			l._logFile.Close()

			l._fullPath = logfile

			l._openLog("02")

			l._writeLog(flag, "--------------------Second Init LogFile--------------------")

		}
	}

	l._writeLog(flag, content)

}

func (l *LogUtils) Close() {
	l._logFile.Close()
}

func (l *LogUtils) Debug(content string) {
	l._appendContent("DEBUG", content)

}
func (l *LogUtils) Info(content string) {
	l._appendContent("INFO", content)

}
func (l *LogUtils) Warn(content string) {
	l._appendContent("WARN", content)

}
func (l *LogUtils) Error(content string) {
	l._appendContent("ERR", content)
}
