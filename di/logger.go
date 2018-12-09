package di

import (
	"memo_sample/adapter/logger"
	"memo_sample/infra/logger"
)

// InjectLog inject log
func InjectLog() logger.Logger {
	return loggersub.NewLogger()
}
