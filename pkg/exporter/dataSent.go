package exporter

import "time"

type DataSent struct {
	Origin string    `json:"origin"`
	Data   []LogSent `json:"data"`
}

type LogSent struct {
	Count     int       `json:"count"`
	FirstSeen time.Time `json:"firstSeen"`
	LastSeen  time.Time `json:"lastSeen"`
	Message   string    `json:"message"`
}
