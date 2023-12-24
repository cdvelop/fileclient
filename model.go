package fileclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type fileCLient struct {
	model.Logger
	model.ObjectHandlerAdapter
	model.DataBaseAdapter
	model.DomAdapter
	model.ThemeAdapter
	model.FetchAdapter

	stringVar string
	err       string

	formData js.Value
}
