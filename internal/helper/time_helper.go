package helper

import "time"

const TimeFormatGmt7 = "2006-01-02T15:04:05.000-07.00"

func FormatTime(t time.Time) string {
	location := time.FixedZone("WIB", 7*60*60)
	return t.In(location).Format(TimeFormatGmt7)
}

func ParseTime(value string) (time.Time, error) {
	return time.Parse(TimeFormatGmt7, value)
}
