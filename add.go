package fileclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func AddFileApi(h *model.MainHandler) {

	fc := &fileCLient{
		Logger:               h.Logger,
		ObjectHandlerAdapter: h.ObjectHandlerAdapter,
		DataBaseAdapter:      h.DataBaseAdapter,
		DomAdapter:           h.DomAdapter,
		ThemeAdapter:         h.ThemeAdapter,
		FetchAdapter:         h.FetchAdapter,
	}

	h.FileApi = fc

	h.FileDiskRW = fc

	js.Global().Set("saveBlobFile", js.FuncOf(fc.saveBlobFile))

}
