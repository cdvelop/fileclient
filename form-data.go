package fileclient

import (
	"syscall/js"
)

func (f *fileCLient) buildFormData(data map[string]any) {

	if blob, ok := data["blob"]; ok {
		// Crear un objeto FormData vacío
		f.formData = js.Global().Get("FormData").New()

		// Agregar el campo y el valor al objeto FormData
		f.formData.Call("append", "id_file", data["id_file"])
		f.formData.Call("append", "object_id", data["object_id"])

		// Extrae el type del Blob
		// blobTypeValue := blob.(js.Value).Get("type").String()
		// Convierte el Blob a un File
		// file := js.Global().Get("File").New([]interface{}{blob}, data["id_file"].(string), map[string]interface{}{"type": blobTypeValue})

		// f.Log("BLOB OK", blob)
		// Agregar el Blob al FormData con un nombre de campo único
		f.formData.Call("append", "files", blob, data["id_file"].(string))
		// f.formData.Call("append", "files", file)

	}

}
