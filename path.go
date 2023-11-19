package fileclient

import (
	"github.com/cdvelop/model"
)

func (f fileCLient) FilePath(params map[string]string) (file_path, file_area string, err error) {
	// fmt.Println("parámetros FilePath recibidos: ", params)

	return file_path, file_area, model.Error("FilePath No implementado en fileclient")
}
