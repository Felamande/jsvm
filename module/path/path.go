package path

import (

	// "errors"

	"path/filepath"

	"github.com/Felamande/jsvm"
	"github.com/Felamande/otto"
)

func init() {
	p := jsvm.Module("path")
	p.Extend("join", join)
	p.Extend("dir", dir)
}

func join(call otto.FunctionCall) otto.Value {
	var pathList []string

	for _, arg := range call.ArgumentList {
		v, _ := arg.Export()
		switch val := v.(type) {
		case []string:
			pathList = append(pathList, val...)
		case string:
			pathList = append(pathList, val)
		}
	}
	path := filepath.Join(pathList...)
	v, _ := otto.ToValue(path)
	return v
}

func dir(call otto.FunctionCall) otto.Value {
	p := call.Argument(0).String()
	return jsvm.StringValue(filepath.Dir(p))
}
