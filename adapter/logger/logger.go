package loggersub

import (
	"log"
	"memo_sample/infra/logger"
)

// NewLogger new log manager
func NewLogger() logger.Logger {
	return lggr{}
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

type lggr struct{}

func (l lggr) Errorf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixError)
	log.Fatalf(format, args...)
}

func (l lggr) Warnf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixWarn)
	log.Fatalf(format, args...)
}

func (l lggr) Infof(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixInfo)
	log.Printf(format, args...)
}

func (l lggr) Debugf(format string, args ...interface{}) {
	log.SetPrefix(LogPrefixDebug)
	log.Printf(format, args...)
}
