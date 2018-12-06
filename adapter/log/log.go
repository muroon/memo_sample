package log

import (
	"log"

	"memo_sample/infra"
)

// NewLog new log manager
func NewLog() infra.Log {
	return logm{}
}

const (
	// LogPrefixError error prefix
	LogPrefixError = "[Error]"

	// LogPrefixWarn warn prefix
	LogPrefixWarn = "[Warnning]"

	// LogPrefixInfo info prefix
	LogPrefixInfo = "[Info]"

	// LogPrefixDebug debug prefix
	LogPrefixDebug = "[Debug]"
)

type logm struct{}

func (l logm) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Fatalf(format, args...)
}

func (l logm) Warnf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixWarn)
	log.Fatalf(format, args...)
}

func (l logm) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}

func (l logm) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
