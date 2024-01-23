package fileclient

import (
	"syscall/js"
)

func (f *fileCLient) saveBlobFile(this js.Value, p []js.Value) any {
	const e = ". saveBlobFile error"
	if len(p) != 3 {
		return f.Log("error 3 arg expected: object name, blob file and extension")
	}

	// f.Log("OBTENEMOS EL NOMBRE DEL OBJETO FILE:", object_file_name)

	f.object, f.err = f.GetObjectBY(p[0].String(), "") //  arg 1
	if f.err != "" {
		return f.Log(f.err + e)
	}

	obj_destiny, err := f.object.Module.GetActualObject()
	if err != "" {
		return f.Log(err + e + " al obtener objeto destino")
	}

	// f.Log("parent_object_name", parent_object_name)

	// OBTENEMOS EL ID DEL OBJETO ACTUAL
	f.stringVar, f.err = obj_destiny.GetID()
	if f.err != "" {
		return f.Log(f.err + e)
	}

	// f.Log("object_id:", f.stringVar)
	id_file, err := f.GetNewID() //id nuevo
	if err != "" {
		return err + e
	}

	var data = map[string]any{
		"id_file":   id_file,
		"object_id": f.stringVar,
		"blob":      p[1], // ARCHIVO BLOB
		"extension": p[2].String(),
	}

	f.err = f.CreateObjectsInDB(f.object.Table, false, data)
	if f.err != "" {
		return f.Log(f.err + e)
	}

	f.buildFormData(data)

	// ENVIAMOS LA DATA AL SERVIDOR PARA SU RESPALDO
	f.SendOneRequest("POST", "upload", f.object.ObjectName, f.formData, func(result []map[string]string, err string) {

		if err != "" {
			f.Log(err + e)
			return
		}
		// f.Log("RESPUESTA upload SERVIDOR:", result)
	})

	if f.object.FrontHandler.ViewHandlerObject != nil {

		html := f.object.FrontHandler.BuildItemsView(map[string]string{
			"id_file":   data["id_file"].(string),
			"url":       data["url"].(string),
			"extension": data["extension"].(string),
		})

		f.err = f.InsertAfterBegin(f.QuerySelectorObject(f.object.ModuleName, f.object.ObjectName), html)
		if f.err != "" {
			f.Log(f.err + e)
		}

	}

	f.new_files_created = true

	return nil

}
