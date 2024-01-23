package fileclient

import (
	"syscall/js"
)

func (f *fileCLient) endFileCapture(this js.Value, p []js.Value) any {

	// f.Log("CAPTURA FINALIZADA")

	if !f.new_files_created {
		return nil
	}

	const e = "endFileCapture error "

	if len(p) != 1 {
		return f.Log(e + "object name arg expected")
	}

	f.object, f.err = f.GetObjectBY(p[0].String(), "") //  arg 1
	if f.err != "" {
		return f.Log(e + f.err)
	}
	// f.Log("OBTENEMOS EL NOMBRE DEL OBJETO FILE:", f.object.ObjectName)

	f.object, f.err = f.object.Module.GetActualObject()
	if f.err != "" {
		return f.Log(e + f.err + " al obtener objeto destino")
	}

	// f.Log("ACTUALIZAMOS LA VISTA OBJETO ACTUAL", f.object.ObjectName)

	if f.object.FrontHandler.AfterUpdate != nil {
		f.err = f.object.FrontHandler.SetObjectInDomAfterUpdate()
		if f.err != "" {
			return f.Log(e + f.err + " SetObjectInDomAfterUpdate")
		}
	}

	f.new_files_created = false

	return nil

}
