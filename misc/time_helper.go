package misc

import (
	"errors"
	"strings"
	"time"
)

const TimeLayoutSimple = "20060102150405"
const TimeLayoutDefault = "2006-01-02 15:04:05"
const TimeLayoutCnShort = "2006/1/2 15:04:05"
const TimeLayoutCnLong = "2006年1月2日 15:04:05"
const TimeLayoutDateSimple = "20060102"
const TimeLayoutDateDefault = "2006-01-02"
const TimeLayoutDateCnShort = "2006/1/2"
const TimeLayoutDateCnLong = "2006年1月2日"

func Now() time.Time {
	return time.Now()
}

func NowInNanos() int64 {
	now := Now()
	return now.UnixNano()
}

func NowInMillis() int64 {
	now := Now()
	unixNano := now.UnixNano()
	return unixNano / 1e6
}

func NowInSeconds() int64 {
	now := Now()
	unixNano := now.UnixNano()
	return unixNano / 1e9
}

func NowInString() string {
	now := Now()
	return now.Format(TimeLayoutDefault)
}

func NowInSimpleString() string {
	now := Now()
	return now.Format(TimeLayoutSimple)
}

func NowDate() time.Time {
	now := Now()
	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	return nowDate
}

func NowDateInMillis() int64 {
	nowDate := NowDate()
	return nowDate.UnixNano() / 1e6
}

func NowDateInSeconds() int64 {
	nowDate := NowDate()
	return nowDate.UnixNano() / 1e9
}

func NowDateInString() string {
	nowDate := NowDate()
	return nowDate.Format(TimeLayoutDateDefault)
}

func NowDateInSimpleString() string {
	nowDate := NowDate()
	return nowDate.Format(TimeLayoutDateSimple)
}

func ParseTime(s string) (time.Time, error) {
	var t time.Time
	var err error
	t, err = time.ParseInLocation(TimeLayoutDefault, s, time.Local)
	if err == nil {
		return t, err
	}
	t, err = time.ParseInLocation(TimeLayoutCnShort, s, time.Local)
	if err == nil {
		return t, err
	}
	t, err = time.ParseInLocation(TimeLayoutCnLong, s, time.Local)
	if err == nil {
		return t, err
	}
	t, err = time.ParseInLocation(TimeLayoutDateDefault, s, time.Local)
	if err == nil {
		return t, err
	}
	t, err = time.ParseInLocation(TimeLayoutDateCnShort, s, time.Local)
	if err == nil {
		return t, err
	}
	t, err = time.ParseInLocation(TimeLayoutDateCnLong, s, time.Local)
	if err == nil {
		return t, err
	}
	return t, errors.New("unknown time format")
}

func FormatTime(t time.Time) string {
	return t.Format(TimeLayoutDefault)
}

func FormatTimeCNShort(t time.Time) string {
	return t.Format(TimeLayoutCnShort)
}

func FormatTimeCNLong(t time.Time) string {
	return t.Format(TimeLayoutCnLong)
}

func FormatDate(t time.Time) string {
	return t.Format(TimeLayoutDateDefault)
}

func FormatDateCNShort(t time.Time) string {
	return t.Format(TimeLayoutDateCnShort)
}

func FormatDateCNLong(t time.Time) string {
	return t.Format(TimeLayoutDateCnLong)
}

func ParseDate(date string) (time.Time, error) {
	filteredDateString := ""
	for _, char := range date {
		if char >= '0' && char <= '9' {
			filteredDateString += string(char)
		} else if len(filteredDateString) > 0 && !strings.HasSuffix(filteredDateString, " ") {
			filteredDateString += " "
		}
	}

	numericDateString := ""
	if strings.Contains(filteredDateString, " ") {
		parts := []string{}
		for _, part := range strings.Split(filteredDateString, " ") {
			if len(part) > 0 {
				parts = append(parts, part)
			}
		}
		if len(parts) >= 1 {
			if len(parts[0]) >= 8 {
				numericDateString += parts[0][0:8]
			} else if len(parts[0]) >= 4 {
				numericDateString += parts[0][0:4]
			} else if len(parts[0]) == 3 {
				numericDateString += "2" + parts[0]
			} else if len(parts[0]) == 2 {
				numericDateString += "20" + parts[0]
			} else if len(parts[0]) == 1 {
				numericDateString += "201" + parts[0]
			} else {
				numericDateString += "1970"
			}
		}
		if len(parts) >= 2 {
			if len(parts[1]) >= 2 {
				numericDateString += parts[1][0:2]
			} else {
				numericDateString += "0" + parts[1]
			}
		}
		if len(parts) >= 3 {
			if len(parts[2]) >= 2 {
				numericDateString += parts[2][0:2]
			} else {
				numericDateString += "0" + parts[2]
			}
		}
	} else {
		numericDateString = filteredDateString
	}
	if len(numericDateString) >= 8 {
		return time.Parse("20060102", numericDateString[0:8])
	} else if len(numericDateString) >= 6 {
		return time.Parse("200601", numericDateString[0:6])
	} else if len(numericDateString) >= 4 {
		return time.Parse("2006", numericDateString[0:4])
	} else if len(numericDateString) == 2 {
		return time.Parse("2006", "20"+numericDateString)
	}
	return time.Parse("20060102", numericDateString)
}

func CompareDate(a time.Time, b time.Time) int {
	if a.Year() != b.Year() {
		return a.Year() - b.Year()
	}
	if a.Month() != b.Month() {
		return int(a.Month()) - int(b.Month())
	}
	return a.Day() - b.Day()
}

func CompareClock(a time.Time, b time.Time) int {
	if a.Hour() != b.Hour() {
		return a.Hour() - b.Hour()
	}
	if a.Minute() != b.Minute() {
		return a.Minute() - b.Minute()
	}
	if a.Second() != b.Second() {
		return a.Second() - b.Second()
	}
	return a.Nanosecond() - b.Nanosecond()
}

func ClockToSecond(clock string) int64 {
	clockInfo := strings.Split(clock, ":")
	var hour int64 = 0
	if len(clockInfo) >= 1 {
		hour = TryParseInt64(clockInfo[0], hour)
	}
	var minute int64 = 0
	if len(clockInfo) >= 2 {
		minute = TryParseInt64(clockInfo[1], minute)
	}
	var second int64 = 0
	if len(clockInfo) >= 3 {
		second = TryParseInt64(clockInfo[2], minute)
	}
	return hour*60*60 + minute*60 + second
}

func NowInClockSpan(beginClock string, endClock string) bool {
	nowInSeconds := NowInSeconds()
	nowDateInSeconds := NowDateInSeconds()
	nowSecond := nowInSeconds - nowDateInSeconds
	startSecond := ClockToSecond(beginClock)
	endSecond := ClockToSecond(endClock)
	return nowSecond > startSecond && nowSecond < endSecond
}
