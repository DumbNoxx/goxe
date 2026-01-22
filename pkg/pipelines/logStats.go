package pipelines

import "time"

type LogStats struct {
	Count     int
	FirstSeen time.Time
	LastSeen  time.Time
	Level     string
}
