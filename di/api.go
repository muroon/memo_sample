package di

import (
	"memo_sample/interface/api"
	"memo_sample/infra"
)

// InjectMemoryAPI inject api
func InjectMemoryAPI() api.API {
	repo := InjectMemoInMemoryRepository() // In Memory
	return api.NewAPI(InjectMemoUsecase(repo))
}

// InjectDBAPI inject api
func InjectDBAPI() api.API {
	repo := InjectMemoDBRepository(infra.NewInfraInfo()) // DB
	return api.NewAPI(InjectMemoUsecase(repo))
}