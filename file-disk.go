package fileclient

import "github.com/cdvelop/model"

func (f fileCLient) FileGet(url_object string) (any, error) {

	return nil, model.Error("No implementado FileGet lectura de archivos en el cliente")
}

func (f fileCLient) FileDelete(path string) error {

	// Borrar archivos desde hdd

	return model.Error("No implementado FileDelete en el cliente")
}
