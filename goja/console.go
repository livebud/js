package goja

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/livebud/js"
)

func console(console *js.Console) map[string]interface{} {
	return map[string]interface{}{
		"log": func(args ...goja.Value) {
			var params []interface{}
			for _, arg := range args {
				params = append(params, arg.String())
			}
			console.Log.Write([]byte(fmt.Sprintln(params...)))
		},
		"warn": func(args ...goja.Value) {
			var params []interface{}
			for _, arg := range args {
				params = append(params, arg.String())
			}
			console.Error.Write([]byte(fmt.Sprintln(params...)))
		},
		"error": func(args ...goja.Value) {
			var params []interface{}
			for _, arg := range args {
				params = append(params, arg.String())
			}
			console.Error.Write([]byte(fmt.Sprintln(params...)))
		},
	}
}
