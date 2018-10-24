package log

import "os"
import "fmt"

type FileAppender struct {
	Appender
	File	*os.File
}

// TODO: rollover
func NewFileAppender(file string, rolloverSize int64) *FileAppender {
	f, err := os.OpenFile(file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &FileAppender{
		File: f,
	}
}

func (c *FileAppender) Log(event *LoggingEvent) {
	c.File.WriteString(fmt.Sprintln(basicLayout(event)))
}

func (c *FileAppender) GetName() string {
	return "file"
}
