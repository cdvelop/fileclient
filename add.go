package fileclient

import (
	"syscall/js"

	"github.com/cdvelop/model"
)

func AddFileApi(h *model.MainHandler) *fileCLient {

	fc := &fileCLient{
		Logger:                h.Logger,
		ObjectsHandlerAdapter: h,
		DataBaseAdapter:       h.DataBaseAdapter,
		DomAdapter:            h.DomAdapter,
		ThemeAdapter:          h.ThemeAdapter,
		FetchAdapter:          h.FetchAdapter,
	}

	h.FileApi = fc

	h.FileDiskRW = fc

	js.Global().Set("saveBlobFile", js.FuncOf(fc.saveBlobFile))
	js.Global().Set("endFileCapture", js.FuncOf(fc.endFileCapture))

	return fc
}
