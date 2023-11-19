package fileclient

import (
	"github.com/cdvelop/model"
)

func (f *fileCLient) FileUpload(object_name, area_file string, request ...any) ([]map[string]string, error) {
	data_out := []map[string]string{}

	return data_out, model.Error("FileUpload No implementado en fileclient")
}
