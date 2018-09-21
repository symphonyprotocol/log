package log

import "time"

type LoggingEvent struct {
	StartTime	time.Time
	Category	string
	Level		Level
	Data		[]interface{}
}

func NewLoggingEvent(category string, level Level, data []interface{}) *LoggingEvent {
	return &LoggingEvent{
		time.Now(),
		category,
		level,
		data,
	}
}
