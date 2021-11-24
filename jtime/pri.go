package jtime

import (
	"strings"
	"time"
)

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
