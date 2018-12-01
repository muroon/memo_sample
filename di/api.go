package di

import (
	"memo_sample/infra"
	"memo_sample/interface/api"
)

// InjectMemoryAPI inject api
func InjectMemoryAPI() api.API {
	repo := InjectMemoInMemoryRepository()
	return api.NewAPI(InjectMemoUsecase(repo))
}

// InjectDBAPI inject api
func InjectDBAPI() api.API {
	repo := InjectMemoDBRepository(infra.NewDBInfo())
	return api.NewAPI(InjectMemoUsecase(repo))
}
