package fileclient

import "github.com/cdvelop/model"

func (f FileCLient) ReadFile(url_object string) ([]byte, error) {

	return nil, model.Error("No implementado ReadFile lectura de archivos en el cliente")
}

func (f FileCLient) DeleteFile(path string) error {

	// Borrar archivos desde hdd

	return model.Error("No implementado DeleteFile en el cliente")
}
