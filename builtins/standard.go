package builtins

import (
	"github.com/brendantang/naiveconcat/data"
)

// StandardDictionary returns a data.Dictionary with bindings for the standard primitive words.
func StandardDictionary() *data.Dictionary {
	return data.NewDictionary(
		map[string]data.Value{
			"say":    data.NewProc(say),
			"+":      data.NewProc(Add),
			"-":      data.NewProc(subtract),
			"*":      data.NewProc(multiply),
			"/":      data.NewProc(divide),
			"define": data.NewProc(define),
			"dup":    data.NewProc(dup),
		},
	)
}
