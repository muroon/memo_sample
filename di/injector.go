//+build wireinject

package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/error"
	"memo_sample/adapter/logger"
	"memo_sample/adapter/memory"
	view "memo_sample/adapter/view/render"
	"memo_sample/interface/api"
	"memo_sample/usecase"

	"github.com/google/wire"
)

// ProvideAPI inject api using wire
var ProvideAPI = wire.NewSet(
	ProvideUsecaseIterator,
	api.NewAPI,
)

// ProvidePresenter inject presenter using wire
var ProvidePresenter = wire.NewSet(
	ProvideRender,
	ProvideLog,
	api.NewPresenter,
	ProvideErrorManager,
)

// ProvideMemoUsecase inject memo usecase using wire
var ProvideMemoUsecase = wire.NewSet(
	ProvideDBRepository, // or ProvideInMemoryRepository
	usecase.NewMemo,
)

// ProvideUsecaseIterator inject usecase itetator using wire
var ProvideUsecaseIterator = wire.NewSet(
	ProvidePresenter,
	ProvideMemoUsecase,
	usecase.NewInteractor,
)

// ProvideInMemoryRepository inject repository using wire
var ProvideInMemoryRepository = wire.NewSet(
	memory.NewTransactionRepository,
	memory.NewMemoRepository,
	memory.NewTagRepository,
)

// ProvideDBRepository inject repository using wire
var ProvideDBRepository = wire.NewSet(
	db.NewTransactionRepository,
	db.NewMemoRepository,
	db.NewTagRepository,
)

// ProvideLog inject log using wire
var ProvideLog = wire.NewSet(loggersub.NewLogger)

// ProvideRender inject render using wire
var ProvideRender = wire.NewSet(view.NewJSONRender)

// ProvideErrorManager inject error manager using wire
var ProvideErrorManager = wire.NewSet(apperrorsub.NewErrorManager)

// InjectAPIServer build inject api using wire
func InjectAPIServer() api.API {
	wire.Build(
		ProvideAPI,
	)
	return nil
}
