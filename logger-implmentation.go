package advanced

import (
	logging "github.com/codemodify/systemkit-logging"
)

type defaultLoggerImplementation struct {
	logger LoggerWithTagAndLevel
}

// NewDefaultLoggerImplementation -
func NewDefaultLoggerImplementation(logger LoggerWithTagAndLevel) LoggerWithTagAndLevelImplementation {
	return &defaultLoggerImplementation{
		logger: logger,
	}
}

func (thisRef defaultLoggerImplementation) TraceWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeTrace))
}

func (thisRef defaultLoggerImplementation) PanicWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypePanic))
}

func (thisRef defaultLoggerImplementation) FatalWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeFatal))
}

func (thisRef defaultLoggerImplementation) ErrorWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeError))
}

func (thisRef defaultLoggerImplementation) WarningWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeWarning))
}

func (thisRef defaultLoggerImplementation) InfoWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeInfo))
}

func (thisRef defaultLoggerImplementation) SuccessWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeSuccess))
}

func (thisRef defaultLoggerImplementation) DebugWithTagAndLevel(tag string, level int, message string) {
	thisRef.logger.Log(NewLogEntryWithTagAndLevel(tag, level, message, logging.TypeDebug))
}
