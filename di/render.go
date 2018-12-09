package di

import (
	"memo_sample/adapter/view/render"
	"memo_sample/view/render"
)

// InjectRender inject render
func InjectRender() render.JSONRender {
	return view.NewJSONRender()
}
