package log

import (
)

// not configurable now...

type Appender interface {
	Log(event *LoggingEvent)
	GetName() string
}

type IAppenderProvider interface {
	log(event *LoggingEvent)
}

type BaseAppender struct {
	EventQueue	chan *LoggingEvent
	AppenderProvider	IAppenderProvider
}

func NewBaseAppender() *BaseAppender {
	b := &BaseAppender{
		EventQueue:	make(chan(*LoggingEvent), 100),
	}
	go b.loop()
	return b
}

func (b *BaseAppender) Log(event *LoggingEvent) {
	b.EventQueue <- event
}

func (b *BaseAppender) loop() {
	for {
		select {
		case event := <- b.EventQueue:
			if b.AppenderProvider != nil {
				b.AppenderProvider.log(event)
			}
		}
	}
}
