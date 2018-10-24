package log

import "fmt"

type ConsoleAppender struct {
	Appender
}

func (c *ConsoleAppender) Log(event *LoggingEvent) {
	fmt.Println(colorizedLayout(event))
}

func (c *ConsoleAppender) GetName() string {
	return "console"
}
