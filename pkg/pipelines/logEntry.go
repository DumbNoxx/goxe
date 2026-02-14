package pipelines

import (
	"sync"
	"time"
)

type LogEntry struct {
	Source    string
	Content   []byte
	Timestamp time.Time
	Level     string
	IdLog     string
	RawEntry  []byte
}

var EntryPool = sync.Pool{
	New: func() any { return new(LogEntry) },
}

var BufferPool = sync.Pool{
	New: func() any {
		return make([]byte, 1024)
	},
}
