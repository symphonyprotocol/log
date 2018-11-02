package log

import "os"
import "fmt"

type FileAppender struct {
	*BaseAppender
}

type FileAppenderProvider struct {
	IAppenderProvider
	File	*os.File
	RolloverSize	int64
}

// TODO: rollover
func NewFileAppender(file string, rolloverSize int64) *FileAppender {
	f, err := os.OpenFile(file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	fa := &FileAppender{
		BaseAppender: NewBaseAppender(),
	}
	fa.AppenderProvider = &FileAppenderProvider{
		File: f,
		RolloverSize: rolloverSize,
	}

	return fa
}

func (c *FileAppenderProvider) log(event *LoggingEvent) {
	c.File.WriteString(fmt.Sprintln(basicLayout(event)))
}

func (c *FileAppender) GetName() string {
	return "file"
}
