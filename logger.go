package log

type Logger struct {
	Name 	string
	Level	Level
}

func NewLogger(category string) *Logger {
	return &Logger{
		category,
		GetLevel(GetLevelForCategory(category), TRACE),
	}
}

func (l *Logger) GetLevel() Level {
	return GetLevel(GetLevelForCategory(l.Name), TRACE)
}

func (l *Logger) SetLevel(level Level) {
	SetLevelForCategory(l.Name, level)
}

func (l *Logger) Log(level Level, params ...interface{}) {
	if IsLevelEnabled(l.GetLevel(), GetLevel(level, INFO)) && IsLevelEnabled(globalLogLevel, GetLevel(level, INFO)) {
		l.log(level, params...)
	}
}

func (l *Logger) log(level Level, params ...interface{}) {
	event := NewLoggingEvent(l.Name, level, params)
	// notify appenders, ugly.

	appenders := GetAppendersForCategory(l.Name)
	if appenders != nil {
		for _, appender := range appenders {
			appender.Log(event)
		}
	}
}

func (l *Logger) Trace(format string, params ...interface{}) {
	l.Log(TRACE, append([]interface{}{ format }, params...)...)
}

func (l *Logger) Debug(format string, params ...interface{}) {
	l.Log(DEBUG, append([]interface{}{ format }, params...)...)
}

func (l *Logger) Info(format string, params ...interface{}) {
	l.Log(INFO, append([]interface{}{ format }, params...)...)
}

func (l *Logger) Warn(format string, params ...interface{}) {
	l.Log(WARN, append([]interface{}{ format }, params...)...)
}

func (l *Logger) Error(format string, params ...interface{}) {
	l.Log(ERROR, append([]interface{}{ format }, params...)...)
}

func (l *Logger) Fatal(format string, params ...interface{}) {
	l.Log(FATAL, append([]interface{}{ format }, params...)...)
}
