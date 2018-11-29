package log

import "os"
import "fmt"

type FileAppender struct {
	*BaseAppender
}

type FileAppenderProvider struct {
	IAppenderProvider
	File	*os.File
	filePath	string
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
		filePath: file,
	}

	return fa
}

func (c *FileAppenderProvider) log(event *LoggingEvent) {
	fi, err := c.File.Stat()
	if err == nil && fi.Size() >= c.RolloverSize {
		// rollover
		
		// try the filename exist
		index := 1
		for ; ; index++ {
			path := fmt.Sprintf("%v.%v", c.filePath, index)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				// not exist
				c.File.Close()
				os.Rename(c.filePath, path)
				
				f, err := os.OpenFile(c.filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}

				c.File = f
				break
			}
		}
	}
	c.File.WriteString(fmt.Sprintln(basicLayout(event)))
}

func (c *FileAppender) GetName() string {
	return "file"
}
