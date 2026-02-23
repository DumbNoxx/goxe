package pipelines

import "time"

type LogBurst struct {
	Category      string
	WindowStart   time.Time
	Count         int
	Ip            string
	AlertsSent    int
	LastAlertTime time.Time
}
