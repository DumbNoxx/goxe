package pipelines

import "time"

type LogStats struct {
	Count    int
	LastSeen time.Time
	Level    string
}
