package advanced

import (
	"time"

	logging "github.com/codemodify/systemkit-logging"
)

// // LoggerWithFields -
// type LoggerWithFields interface {
// 	logging.Logger

// 	TraceWithFields(fields logging.Fields)
// 	PanicWithFields(fields logging.Fields)
// 	FatalWithFields(fields logging.Fields)
// 	ErrorWithFields(fields logging.Fields)
// 	WarningWithFields(fields logging.Fields)
// 	InfoWithFields(fields logging.Fields)
// 	SuccessWithFields(fields logging.Fields)
// 	DebugWithFields(fields logging.Fields)
// }

// LogEntryWithTagAndLevel -
type LogEntryWithTagAndLevel struct {
	logging.LogEntry        //
	Tag              string // "This-Is-A-Tag"
	Level            int    // Ex: parentMethod - level 0, childMethod() - level 1, useful for concurrent sorted logging with call-stack alike
}

// NewLogEntryWithTagAndLevel -
func NewLogEntryWithTagAndLevel(tag string, level int, message string, logType logging.LogType) LogEntryWithTagAndLevel {
	r := LogEntryWithTagAndLevel{
		Tag:   tag,
		Level: 0,
	}
	r.Time = time.Now()
	r.Type = logType
	r.Message = message

	return r
}

// LoggerWithTagAndLevel -
type LoggerWithTagAndLevel interface {
	Log(logEntry LogEntryWithTagAndLevel) LogEntryWithTagAndLevel
}

// LoggerWithTagAndLevelImplementation -
type LoggerWithTagAndLevelImplementation interface {
	TraceWithTagAndLevel(tag string, level int, message string)
	PanicWithTagAndLevel(tag string, level int, message string)
	FatalWithTagAndLevel(tag string, level int, message string)
	ErrorWithTagAndLevel(tag string, level int, message string)
	WarningWithTagAndLevel(tag string, level int, message string)
	InfoWithTagAndLevel(tag string, level int, message string)
	SuccessWithTagAndLevel(tag string, level int, message string)
	DebugWithTagAndLevel(tag string, level int, message string)
}
