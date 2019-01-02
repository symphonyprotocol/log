package log

import "fmt"

type ConsoleAppender struct {
	*BaseAppender
}

type ConsoleAppenderProvider struct {
	IAppenderProvider
}

func NewConsoleAppender() *ConsoleAppender {
	c := &ConsoleAppender{
		BaseAppender: NewBaseAppender(),
	}
	c.AppenderProvider = &ConsoleAppenderProvider{}
	return c
}

func (cp *ConsoleAppenderProvider) log(event *LoggingEvent) {
	fmt.Println(colorizedLayout(event))
}

func (c *ConsoleAppender) GetName() string {
	return "console"
}
