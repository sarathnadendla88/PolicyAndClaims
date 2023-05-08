package date_utils

import (
	"time"
)

var dateFormatYYYYMMDDTHHMMSSZ = "2006-01-02T15:04:05Z"
var dateFormatYYYYMMDDHHMMSS = "2006-01-02 15:04:05"

func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	return time.Now().In(loc)
}

func GetNowString() string {
	return GetNowTime().Format(dateFormatYYYYMMDDTHHMMSSZ)
}

func GetNowStringForDB() string {
	return GetNowTime().Format(dateFormatYYYYMMDDHHMMSS)
}
