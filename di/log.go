package di

import (
	"memo_sample/adapter/log"
	"memo_sample/infra"
)

// InjectLog inject log
func InjectLog() infra.Log {
	return log.NewLog()
}
