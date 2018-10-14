package di

import (
	"memo_sample/interface/api"
)

// InjectAPI inject api
func InjectAPI() api.API {
	repo := InjectMemoInMemoryRepository() // In Memory
	return api.NewAPI(InjectMemoUsecase(repo))
}