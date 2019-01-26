package log

import "fmt"
import "runtime"

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
	if runtime.GOOS == "windows" {
		fmt.Fprintln(Output, colorizedLayout(event))
	} else {
		fmt.Println(colorizedLayout(event))
	}
}

func (c *ConsoleAppender) GetName() string {
	return "console"
}
