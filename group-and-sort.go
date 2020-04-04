package advanced

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// Flushable -
type Flushable interface {
	LoggerWithTagAndLevel

	Flush()
}

type groupAndSortLogger struct {
	logger                LoggerWithTagAndLevel
	logLinesByKey         map[string][]LogEntryWithTagAndLevel
	logLinesByKeyWithTime map[string]time.Time
	logLinesByKeyMutex    *sync.RWMutex
}

// NewGroupAndSort -
func NewGroupAndSort(logger LoggerWithTagAndLevel) Flushable {
	return &groupAndSortLogger{
		logger:                logger,
		logLinesByKey:         map[string][]LogEntryWithTagAndLevel{},
		logLinesByKeyWithTime: map[string]time.Time{},
		logLinesByKeyMutex:    &sync.RWMutex{},
	}
}

func (thisRef *groupAndSortLogger) Log(logEntry LogEntryWithTagAndLevel) LogEntryWithTagAndLevel {
	thisRef.logLinesByKeyMutex.Lock()
	defer thisRef.logLinesByKeyMutex.Unlock()

	thisRef.logLinesByKey[logEntry.Tag] = append(thisRef.logLinesByKey[logEntry.Tag], logEntry)

	// Remember the earliest log-line
	var ok bool
	if _, ok = thisRef.logLinesByKeyWithTime[logEntry.Tag]; !ok {
		thisRef.logLinesByKeyWithTime[logEntry.Tag] = logEntry.Time
	} else {
		var storedTime = thisRef.logLinesByKeyWithTime[logEntry.Tag]
		if storedTime.After(logEntry.Time) {
			thisRef.logLinesByKeyWithTime[logEntry.Tag] = logEntry.Time
		}
	}

	return logEntry
}

func (thisRef *groupAndSortLogger) Flush() {
	thisRef.logLinesByKeyMutex.RLock()
	defer thisRef.logLinesByKeyMutex.RUnlock()

	// Sort by time
	var allTimes = []time.Time{}
	var timeToLogEntryTag = map[int64]string{}
	for key, value := range thisRef.logLinesByKeyWithTime {
		allTimes = append(allTimes, value)
		timeToLogEntryTag[value.UnixNano()] = key
	}
	sort.Slice(
		allTimes,
		func(i, j int) bool {
			return allTimes[i].Before(allTimes[j])
		},
	)

	for index := range allTimes {
		var logEntryTag = timeToLogEntryTag[allTimes[index].UnixNano()]
		var arrayOfLogEntries = thisRef.logLinesByKey[logEntryTag]

		sort.Slice(
			arrayOfLogEntries,
			func(i, j int) bool {
				return arrayOfLogEntries[i].Time.Before(arrayOfLogEntries[j].Time)
			},
		)

		for i := range arrayOfLogEntries {
			arrayOfLogEntries[i].Tag = fmt.Sprintf("%10s", arrayOfLogEntries[i].Tag)
			thisRef.logger.Log(arrayOfLogEntries[i])
		}
	}
}
