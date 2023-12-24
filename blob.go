package fileclient

import (
	"syscall/js"
)

func (f fileCLient) saveBlobFile(this js.Value, p []js.Value) any {
	const e = ". saveBlobFile"
	if len(p) != 2 {
		return f.Log("error arg expected: object name, object id destiny and blob file")
	}

	object_file_name := p[0].String() //  arg 1
	// f.Log("OBTENEMOS EL NOMBRE DEL OBJETO FILE:", object_file_name)

	// ARCHIVO BLOB
	blob := p[1] // arg 2

	parent_object_name := f.ObjectActual().ObjectName
	// f.Log("parent_object_name", parent_object_name)

	// OBTENEMOS EL ID DEL OBJETO ACTUAL
	f.stringVar, f.err = f.ObjectActual().GetID()
	if f.err != "" {
		return f.Log(f.err + e)
	}

	// f.Log("object_id:", f.stringVar)

	var data = map[string]any{
		"object_id": f.stringVar,
		"blob":      blob,
	}

	// SELECCIONAMOS EL OBJETO FILE PARA GUARDAR EN DB
	f.err = f.SetActualObject(object_file_name)
	if f.err != "" {
		return f.Log(f.err + e)
	}
	// f.Log("OBJETO FILE PARA GUARDAR:", f.ObjectActual().ObjectName)

	f.err = f.CreateObjectsInDB(f.ObjectActual().Table, true, data)
	if f.err != "" {
		return f.Log(f.err + e)
	}
	f.Log("GUARDADO DATA EN DB LOCAL OK:", data)

	// ENVIAMOS LA DATA AL SERVIDOR PARA SU RESPALDO

	f.buildFormData(data)

	f.SendOneRequest("POST", "upload", object_file_name, f.formData, func(result []map[string]string, err string) {

		if err != "" {
			f.Log(err)
			return
		}

		f.Log("RESPUESTA upload SERVIDOR:", result)

	})

	if f.ObjectActual().FrontHandler.ObjectViewHandler != nil {

		fiel_id := f.ObjectActual().PrimaryKeyName()

		// f.Log("ObjectActual fiel_id:", fiel_id)

		html := f.ObjectActual().FrontHandler.BuildItemsView(map[string]string{
			fiel_id: data[fiel_id].(string),
			"url":   data["url"].(string),
		})

		f.err = f.InsertAfterBegin(f.QuerySelectorObject(f.ObjectActual().ModuleName, f.ObjectActual().ObjectName), html)
		if f.err != "" {
			f.Log(f.err + e)
		}

	}

	// "VOLVEMOS AL OBJETO PADRE"
	f.err = f.SetActualObject(parent_object_name)
	if f.err != "" {
		return f.Log(f.err + e)
	}

	return nil

}
