package advanced

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const maxFormattedTimeLength = 30

type simpleFormatter struct{}

// NewSimpleFormatter -
func NewSimpleFormatter() LoggerWithTagAndLevel {
	return &simpleFormatter{}
}

func (thisRef simpleFormatter) Log(logEntry LogEntryWithTagAndLevel) LogEntryWithTagAndLevel {
	var formattedTime = logEntry.Time.UTC().Format(time.RFC3339Nano)

	// reformat the time and fill-in with zeros for the nano seconds
	if len(formattedTime) < maxFormattedTimeLength {
		var spacesCount = maxFormattedTimeLength - len(formattedTime)

		var newV = fmt.Sprintf("%"+strconv.Itoa(spacesCount+1)+"v", "Z")
		newV = strings.Replace(newV, " ", "0", spacesCount)

		formattedTime = strings.Replace(
			formattedTime,
			"Z",
			newV,
			1,
		)
	}

	// format the log line
	var formatting = "%s | %s"
	if len(strings.TrimSpace(logEntry.Tag)) > 0 {
		formatting = formatting + " | %s"
	} else {
		formatting = formatting + " |"
	}

	if logEntry.Level > 0 {
		formatting = formatting + fmt.Sprintf(" %"+strconv.Itoa(logEntry.Level*4)+"v", "")
		formatting += " ->"
	}
	formatting = formatting + " %s"

	if len(strings.TrimSpace(logEntry.Tag)) > 0 {
		logEntry.Message = fmt.Sprintf(
			formatting,
			formattedTime,
			logEntry.Type,
			logEntry.Tag,
			logEntry.Message,
		)
	} else {
		logEntry.Message = fmt.Sprintf(
			formatting,
			formattedTime,
			logEntry.Type,
			logEntry.Message,
		)
	}

	return logEntry
}
