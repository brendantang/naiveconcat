package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

// Dict returns a data.Dictionary with bindings for the standard primitive words.
func Dict() *data.Dictionary {
	return data.NewDictionary(
		nil,
		map[string]data.Value{
			"say":        data.NewProc(say),
			"show-stack": data.NewProc(showStack),
			"+":          data.NewProc(add),
			"-":          data.NewProc(subtract),
			"*":          data.NewProc(multiply),
			"/":          data.NewProc(divide),
			"=":          data.NewProc(equal),
			"let":        data.NewProc(let),
			"dup":        data.NewProc(dup),
			"true":       data.NewBoolean(true),
			"false":      data.NewBoolean(false),
			"not":        data.NewProc(not),
			"or":         data.NewProc(or),
			"length":     data.NewProc(length),
			"lop":        data.NewProc(lop),
		},
	)
}
