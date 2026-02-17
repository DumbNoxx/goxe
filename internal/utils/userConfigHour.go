package utils

import (
	"log"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
)

var TimeReportFile = UserConfigHour()

// UserConfigHour calculates the duration until the next occurrence of the configured hour
// (Config.GenerateLogsOptions.Hour) on the current or following day.
//
// Returns:
//
//   - time.Duration: The time remaining to reach the target hour. If the hour has already passed today,
//     it returns the duration until that same hour on the following day.
//
// The function performs:
//
//   - Retrieves the hour string from options.Config.GenerateLogsOptions.Hour (expected format: "15:04:05").
//   - Parses the hour using time.Parse. If the format is invalid, the program terminates with log.Fatal.
//   - Constructs a time.Time object for today at the parsed hour, using the current local time zone.
//   - Calculates the difference between that target time and the current moment (time.Now()).
//   - If the difference is negative (the hour has already passed), it adds 24 hours to target the next day.
//   - Returns the resulting duration.
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
