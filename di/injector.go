//+build wireinject

package di

import (
	"memo_sample/adapter/db"
	"memo_sample/adapter/logger"
	"memo_sample/adapter/memory"
	view "memo_sample/adapter/view/render"
	"memo_sample/interface/api"
	"memo_sample/usecase"

	"github.com/google/wire"
)

// WireInjectAPI inject api using wire
var WireInjectAPI = wire.NewSet(
	WireInjectUsecaseIterator,
	api.NewAPI,
)

// WireInjectPresenter inject presenter using wire
var WireInjectPresenter = wire.NewSet(
	WireInjectRender,
	WireInjectLog,
	api.NewPresenter,
)

// WireInjectMemoUsecase inject memo usecase using wire
var WireInjectMemoUsecase = wire.NewSet(
	WireInjectDBRepository, // or WireInjectInMemoryRepository
	usecase.NewMemo,
)

// WireInjectUsecaseIterator inject usecase itetator using wire
var WireInjectUsecaseIterator = wire.NewSet(
	WireInjectPresenter,
	WireInjectMemoUsecase,
	usecase.NewInteractor,
)

// WireInjectInMemoryRepository inject repository using wire
var WireInjectInMemoryRepository = wire.NewSet(
	memory.NewTransactionRepository,
	memory.NewMemoRepository,
	memory.NewTagRepository,
)

// WireInjectDBRepository inject repository using wire
var WireInjectDBRepository = wire.NewSet(
	db.NewTransactionRepository,
	db.NewMemoRepository,
	db.NewTagRepository,
)

// WireInjectLog inject log using wire
var WireInjectLog = wire.NewSet(loggersub.NewLogger)

// WireInjectRender inject render using wire
var WireInjectRender = wire.NewSet(view.NewJSONRender)

// InjectAPIServer build inject api using wire
func InjectAPIServer() api.API {
	wire.Build(
		WireInjectAPI,
	)
	return nil
}
