package pipelines

import "time"

type LogEntry struct {
	Source    string
	Content   string
	Timestamp time.Time
	Level     string
	IdLog     string
}
