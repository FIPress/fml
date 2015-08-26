package fipml

import (
	"strconv"
	"time"
	"regexp"
	. "fiputil"
)

const (
	dateF = "2006-01-02"
	datetimeF = "2006-01-02 15:04:05"
)

var (
	dateR = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	datetimeR = regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	datetimeRFC3399R = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(.\d+)?(Z|[+-]\d{2}:\d{2})$`)
)

func eval(raw string) (val interface {}) {
	var ok bool
	switch raw[0] {
	case 't','T','f','F':
		val,ok = evalBool(raw)
		if ok {
			return
		}
	case '+', '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		val, ok = evalInt(raw)
		if ok {
			return
		}
		val, ok = evalFloat(raw)
		if ok {
			return
		}
		val, ok = evalDatetime(raw)
		if ok {
			return
		}
	}
	//log.Println("raw:",raw)
	return Unquote(raw)
}

func evalString(raw string) string {
	return Unquote(raw)
}

func evalBool(raw string) (val bool,ok bool) {
	switch raw {
	case "t", "T", "true", "TRUE", "True":
		return true, true
	case "f", "F", "false", "FALSE", "False":
		return false, true
	}
	return false, false
}

func evalInt(raw string) (val int, ok bool) {
	val,err := strconv.Atoi(raw)
	return val, err == nil
}

func evalFloat(raw string) (val float64,ok bool) {
	val,err := strconv.ParseFloat(raw,64)
	return val, err == nil
}

func evalDatetime(raw string) (val time.Time, ok bool) {
	var err error
	if dateR.MatchString(raw) {
		val,err = time.Parse(dateF,raw)
	} else if datetimeR.MatchString(raw) {
		val,err = time.Parse(datetimeF,raw)
	} else if datetimeRFC3399R.MatchString(raw) {
		val,err = time.Parse(time.RFC3339,raw)
	} else {
		return MinTime, false
	}
	return val, err == nil
}

func doEvalDatetime(raw string,format string) (val time.Time, ok bool) {
	val, err := time.Parse(format,raw)
	return val, err == nil
}

