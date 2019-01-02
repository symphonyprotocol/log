package log

import "sync"

var loggers sync.Map
var globalLogLevel Level
var configuration sync.Map

func GetLogger(category string) *Logger {
	if _logger, ok := loggers.Load(category); ok {
		return _logger.(*Logger)
	}

	if category == "" {
		category = "default"
	} 
	_logger := NewLogger(category)
	loggers.Store(category, _logger)
	return _logger
}

func GetDefaultLogger() *Logger {
	return GetLogger("")
}

func SetGlobalLevel(level Level) {
	globalLogLevel = level
}

func Configure(config map[string]([]Appender)) {
	if config != nil {
		for k, v := range config {
			configuration.Store(k, v)
		}
	}
}

func AddConfiguration(category string, appenders []Appender) {
	configuration.Store(category, appenders)
}

func GetAppendersForCategory(category string) []Appender {
	if _appender, _ok := configuration.Load(category); _ok {
		if appender, ok := _appender.([]Appender); ok {
			return appender
		}
	} else {
		// use default appender
		if _appender, _ok := configuration.Load("default"); _ok {
			if appender, ok := _appender.([]Appender); ok {
				return appender
			}
		}
	}

	return nil
}
