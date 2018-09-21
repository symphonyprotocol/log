package log

import (
	"strings"
	"math"
)

type Level string

const (
	ALL Level = "all"
	TRACE Level = "trace"
	DEBUG Level = "debug"
	INFO Level = "info"
	WARN Level = "warn"
	ERROR Level = "error"
	FATAL Level = "fatal"
	NONE Level = "none"
)

var (
	LogLevelPriorities = map[Level] uint {
		ALL: 0,
		TRACE: 1000,
		DEBUG: 2000,
		INFO: 3000,
		WARN: 4000,
		ERROR: 5000,
		FATAL: 6000,
		NONE: math.MaxUint32,
	}
	LogLevelColors = map[Level] string {
		ALL: "grey",
		TRACE: "blue",
		DEBUG: "cyan",
		INFO: "green",
		WARN: "yellow",
		ERROR: "red",
		FATAL: "magenta",
		NONE: "grey",
	}
)

func GetLevel(level Level, defaultLevel Level) Level {
	if level != "" {
		return level
	}

	return defaultLevel
}

func IsLevelEnabled(level Level, targetLevel Level) bool {
	return LogLevelPriorities[level] <= LogLevelPriorities[targetLevel]
}

func (l Level) String() string {
	return strings.ToUpper(string(l))
}
