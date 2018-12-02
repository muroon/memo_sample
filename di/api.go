package di

import (
	"memo_sample/infra"
	"memo_sample/interface/api"
)

// InjectMemoryAPI inject api
func InjectMemoryAPI() api.API {
	return api.NewAPI(InjectMemoUsecase(InjectInMemoryRepository()))
}

// InjectDBAPI inject api
func InjectDBAPI() api.API {
	return api.NewAPI(InjectMemoUsecase(InjectDBRepository(infra.NewDBInfo())))
}
