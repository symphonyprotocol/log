package log

// not configurable now...

type Appender interface {
	Log(event *LoggingEvent)
}

var appenders = []Appender {
	&ConsoleAppender{},
}
