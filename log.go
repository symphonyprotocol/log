package log

import "sync"

var loggers sync.Map

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
