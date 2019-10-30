package common

import (
	"fmt"
	"strconv"
	"time"
)

func FormatTimestamp(timestamp int64) (year string, month string, day string, hour string, minute string, second string) {
	tm := time.Unix(timestamp, 0)
	y, m, d := tm.Date()
	hr, min, s := tm.Clock()

	return strconv.Itoa(y), string(m.String()), strconv.Itoa(d), strconv.Itoa(hr), strconv.Itoa(min), strconv.Itoa(s)
}

func FormatTimeExpiredDonation(timestamp int64) (formattedTimestamp string) {
	year, month, day, hour, minute, _ := FormatTimestamp(timestamp)
	stringTime := fmt.Sprintf("%s %s %s %s:%s", day, month, year, hour, minute)

	dt, _ := time.Parse("02 January 2006 15:04", stringTime)
	formattedTimestamp = dt.Format("02 January 2006 15:04 WIB")
	return formattedTimestamp
}

func TranslateDateToIndonesian(date time.Time) string {
	location, _ := time.LoadLocation("Asia/Jakarta")
	localDate := date.In(location)

	weekday := WeekdayId[date.Weekday()]
	day := localDate.Day()
	month := MonthId[localDate.Month()]
	year := localDate.Year()
	hour := localDate.Hour()
	minute := localDate.Minute()
	timeZone, _ := localDate.Zone()

	return fmt.Sprintf("%s, %d %s %d, %02d:%02d %s", weekday, day, month, year, hour, minute, timeZone)
}
