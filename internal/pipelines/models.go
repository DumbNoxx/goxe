package pipelines

import "time"

type LogEntry struct {
	Source    string
	Content   string
	Timestamp time.Time
	Level     string
}

type LogStats struct {
	Count    int
	LastSeen time.Time
	Level    string
}
