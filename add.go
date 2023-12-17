package fileclient

import "github.com/cdvelop/model"

func AddFileApi(h *model.MainHandler) {

	fc := &fileCLient{}

	h.FileApi = fc

	h.FileDiskRW = fc

}
