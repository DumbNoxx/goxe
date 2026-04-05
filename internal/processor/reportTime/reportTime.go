package reporttime

import (
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/internal/utils"
)

// GetReportFileTime stops and restarts the report file generation ticker
// using the interval defined in 'utils.TimeReportFile'.
//
// The function performs:
//
//   - Stops 'processor.TickerReportFile' using Stop()
//   - Restarts it using Reset(utils.TimeReportFile) to apply new interval.
func GetReportFileTime(getConfig options.ConfigProvider) {
	processor.TickerReportFile.Stop()
	processor.TickerReportFile.Reset(utils.UserConfigHour(getConfig))
}

// GetReportPartialTime updates the partial report interval based on the current
// configuration (options.Config.ReportInterval) and restarts the corresponding ticker.
//
// The function performs:
//
//   - Calculates processor.TimeReport as options.Config.ReportInterval (in minutes)
//     converted to time.Duration
//   - Stops processor.Ticker using Stop()
//   - Restarts the ticket with the new interval using Reset(processor.TimeReport),
func GetReportPartialTime(getConfig options.ConfigProvider) {
	conf := getConfig()
	processor.TimeReport = time.Duration(conf.ReportInterval * float64(time.Minute))
	processor.Ticker.Stop()
	processor.Ticker.Reset(processor.TimeReport)
}
