package eval

import (
	"github.com/brendantang/naiveconcat/data"
)

// CoreDict returns a data.Dictionary with word bindings for built-in primitive
// values.
func CoreDict() *data.Dictionary {
	return data.NewDictionary(
		nil,
		map[string]data.Value{
			// IO
			"say":   data.NewProc(say),       // pop a value and print it
			"stack": data.NewProc(showStack), // print the whole stack
			"words": data.NewProc(showDict),  // print all word definitions

			// MATH
			"+": data.NewProc(add),      // pop two numbers and push their sum
			"-": data.NewProc(subtract), // pop a, pop b, push b - a
			"*": data.NewProc(multiply), // pop two numbers and push their product
			"/": data.NewProc(divide),   // pop a, pop b, push b divided by a
			"=": data.NewProc(equal),    // pop two numbers and push whether they're equal

			// BOOLEANS
			"true":  data.NewBoolean(true),  // TRUE literal
			"false": data.NewBoolean(false), // FALSE literal
			"not":   data.NewProc(not),      // pop a boolean, push its negation
			"or":    data.NewProc(or),       // pop two bools, push a bool saying whether either is true
			"and":   data.NewProc(and),      // pop two bools, push a bool saying whether both are true
			"then":  data.NewProc(then),     // pop a bool and x, evaluate x if the bool is TRUE

			// QUOTATIONS
			"length": data.NewProc(length), // push length of the quotation on top of the stack
			"lop":    data.NewProc(lop),    // pop quotation { a b c ... }, push { b c ... }, push a

			// DICTIONARY MANIPULATION
			"let":    data.NewProc(let),    // pop a string and a value, make a word named by string, defined by value
			"define": data.NewProc(define), // pop a string and x. make a word named by string, defined by proc that evals to x.

			// STACK MANIPULATION
			"dup":    data.NewProc(dup),    // pop a, push a, push a
			"drop":   data.NewProc(drop),   // pop a, discard it
			"lambda": data.NewProc(lambda), // pop x, push a proc that evals to x.

			// APPLY
			"apply": data.NewProc(apply), // pop x. if x is a quotation, eval each item. otherwise, eval x.
		},
	)
}
