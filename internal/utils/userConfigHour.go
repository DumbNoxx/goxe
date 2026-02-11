package utils

import (
	"log"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
)

var TimeReportFile = UserConfigHour()

func UserConfigHour() (dateHour time.Duration) {
	userHour := options.Config.GenerateLogsOptions.Hour
	parseHour, err := time.Parse("15:04:05", userHour)
	hourNow := time.Now()

	if err != nil {
		log.Fatal("Field 'Config.GenerateLogsOptions.Hour' (config.json) must be in 00:00:00 format.")
	}

	hourToday := time.Date(hourNow.Year(), hourNow.Month(), hourNow.Day(), parseHour.Hour(), parseHour.Minute(), parseHour.Second(), 0, hourNow.Location())
	dateHourTime := hourToday.Sub(hourNow)
	if dateHourTime < 0 {
		dateHour = dateHourTime + (time.Hour * 24)
	} else {
		dateHour = dateHourTime
	}

	return dateHour
}
