package fileclient

import "github.com/cdvelop/model"

func AddFileApi(h *model.Handlers) {

	fc := &fileCLient{}

	h.FileApi = fc

	h.FileDiskRW = fc

}
