package jlog2

import "github.com/yjhi/golib/jlog"

var LogRoot string = "logs"
var LogName string = "day"
var LogTail string = "log"
var _logUtil *jlog.LogUtils = nil

func _initLog() {
	if _logUtil == nil {
		_logUtil = jlog.BuildDayLogUtils(LogRoot, LogName, LogTail)
	}
}

func LogDebug(text string) {
	_initLog()

	_logUtil.Debug(text)
}

func LogInfo(text string) {
	_initLog()

	_logUtil.Info(text)
}
func LogError(text string) {
	_initLog()

	_logUtil.Error(text)
}

func LogWarn(text string) {
	_initLog()

	_logUtil.Warn(text)
}
