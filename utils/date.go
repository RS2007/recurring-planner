package utils

import (
	"time"
)

func ConvertUnixToDateString(unixTimestamp int64) string {
	timeStruct := time.Unix(unixTimestamp, 0)
	stringDate := timeStruct.Format(time.UnixDate)
	return stringDate
}

func GetFormattedTimeString(timeString string) (string, error) {
	layout := time.RFC3339
	t, err := time.Parse(layout, timeString)
	if err != nil {
		return "", err
	}
	return t.Format("15:04:05 -07:00"), nil
}

func GetDateTimeString(dateString string, timeString string, timeZone string) (string, error) {
	nonStringTime, err := time.Parse("2006-01-02 15:04:05 -07:00", dateString+" "+timeString)

	ErrorHandler(err)
	return nonStringTime.Format("2006-01-02T15:04:05-0700"), nil
}
