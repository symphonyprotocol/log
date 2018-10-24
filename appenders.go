package log

// not configurable now...

type Appender interface {
	Log(event *LoggingEvent)
	GetName() string
}
