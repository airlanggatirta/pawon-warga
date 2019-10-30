package common

import "time"

var (
	WeekdayId map[time.Weekday]string = make(map[time.Weekday]string)
	MonthId   map[time.Month]string   = make(map[time.Month]string)
)

func init() {
	WeekdayId[time.Sunday] = "Minggu"
	WeekdayId[time.Monday] = "Senin"
	WeekdayId[time.Tuesday] = "Selasa"
	WeekdayId[time.Wednesday] = "Rabu"
	WeekdayId[time.Thursday] = "Kamis"
	WeekdayId[time.Friday] = "Jumat"
	WeekdayId[time.Saturday] = "Sabtu"

	MonthId[time.January] = "Januari"
	MonthId[time.February] = "Februari"
	MonthId[time.March] = "Maret"
	MonthId[time.April] = "April"
	MonthId[time.May] = "Mei"
	MonthId[time.June] = "Juni"
	MonthId[time.July] = "Juli"
	MonthId[time.August] = "Agustus"
	MonthId[time.September] = "September"
	MonthId[time.October] = "Oktober"
	MonthId[time.November] = "November"
	MonthId[time.December] = "Desember"
}
