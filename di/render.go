package di

import (
	view "memo_sample/adapter/view/api"
	"memo_sample/interface/api"
)

// InjectRender inject render
func InjectRender() api.APIRender {
	return view.NewAPIRender()
}
