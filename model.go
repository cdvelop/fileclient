package fileclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

type fileCLient struct {
	model.Logger
	model.ObjectsHandlerAdapter
	model.DataBaseAdapter
	model.DomAdapter
	model.ThemeAdapter
	model.FetchAdapter

	stringVar string
	object    *model.Object
	err       string

	formData js.Value

	new_files_created bool
}
