package log

import (
	"fmt"
)

var styles map[string] []int = map[string] []int {
	"bold": { 1, 22 },
	"italic": { 3, 23 },
	"underline": { 4, 24 },
	"inverse": { 7, 27 },

	"white": { 37, 39 },
	"grey": { 90, 39 },
	"black": { 90, 39 },

	"blue": { 34, 39 },
	"cyan": { 36, 39 },
	"green": { 32, 39 },
	"magenta": {35, 39 },
	"red": { 91, 39 },
	"yellow": { 33, 39 },
}

func colorStart(color string) string {
	if color != "" {
		return fmt.Sprintf("\x1B[%vm", styles[color][0])
	}
	return ""
}

func colorEnd(color string) string {
	if color != "" {
		return fmt.Sprintf("\x1B[%vm", styles[color][1])
	}
	return ""
}

func colorize(str string, color string) string {
	return colorStart(color) + str + colorEnd(color)
}

func basicLayout(event *LoggingEvent) string {
	return fmt.Sprintf("[%v] [%v] %v - ", event.StartTime.Format("2006-01-02T15:04:05.999-07:00"), event.Level.String(), event.Category) + fmt.Sprintf(event.Data[0].(string), event.Data[1:]...)
}

func colorizedLayout(event *LoggingEvent) string {
	return colorize(fmt.Sprintf("[%v] [%v] %v - ", event.StartTime.Format("2006-01-02T15:04:05.999-07:00"), event.Level.String(), event.Category), LogLevelColors[event.Level]) + fmt.Sprintf(event.Data[0].(string), event.Data[1:]...)
}

