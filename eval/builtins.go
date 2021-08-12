package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"sort"
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
			"length": data.NewProc(length),       // push length of the quotation on top of the stack
			"lop":    data.NewProc(lop),          // pop quotation { a b c ... }, push { b c ... }, push a
			"append": data.NewProc(appendToQuot), // pop quotation { a b c ... }, pop value d, push { a b c ... d }

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



// IO.


func say(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}
	fmt.Println(val)
	return nil
}


func showStack(d *data.Dictionary, s *data.Stack) error {
	fmt.Println(s)
	return nil
}


func showDict(d *data.Dictionary, s *data.Stack) error {
	words := make([]string, 0, len(d.Bindings))
	for w := range d.Bindings {
		words = append(words, w)
	}
	sort.Strings(words)
	for _, w := range words {
		fmt.Printf("%s\t%s\n", w, d.Bindings[w])
	}
	return nil
}



// Math.


func add(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number + a.Number))

	return nil
}


func subtract(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number - a.Number))

	return nil
}


func multiply(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(a.Number * b.Number))

	return nil
}


func divide(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewNumber(b.Number / a.Number))

	return nil
}


func equal(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Number)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Number == a.Number))

	return nil
}



// Booleans.


// not pops a boolean and pushes its negation.
func not(d *data.Dictionary, s *data.Stack) error {
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(!b.Bool))
	return nil
}


// or pops two booleans and pushes TRUE if either of them is TRUE, FALSE
// otherwise.
func or(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Bool || a.Bool))
	return nil
}


// and pops two booleans and pushes TRUE if both of them are TRUE, FALSE
// otherwise.
func and(d *data.Dictionary, s *data.Stack) error {
	a, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	b, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	s.Push(data.NewBoolean(b.Bool && a.Bool))
	return nil
}


// then pops a predicate and a value. If the predicate is TRUE, push the value.
// Otherwise discard it.
func then(d *data.Dictionary, s *data.Stack) error {
	predicate, err := s.PopType(data.Boolean)
	if err != nil {
		return err
	}
	consequent, err := s.Pop()
	if err != nil {
		return err
	}
	if predicate.Bool {
		err := Eval(consequent, d, s)
		if err != nil {
			return err
		}
	}
	return nil
}



// Quotation manipulation.


// length returns the length of the quotation on top of the stack without
// popping it.
func length(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.Peek()
	if err != nil {
		return err
	}
	if quot.Type != data.Quotation {
		return data.NewTypeErr(quot, data.Quotation)
	}
	s.Push(data.NewNumber(float64(len(quot.Quotation))))

	return nil
}


// lop pops the quotation on top of the stack, then pushes its tail, then pushes
// its head.
func lop(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.PopType(data.Quotation)
	if err != nil {
		return err
	}
	head, tail := quot.Quotation[0], quot.Quotation[1:]
	s.Push(data.NewQuotation(tail...))
	s.Push(head)

	return nil
}


// appendToQuot adds a value to the end of a quotation.
func appendToQuot(d *data.Dictionary, s *data.Stack) error {
	quot, err := s.PopType(data.Quotation)
	if err != nil {
		return err
	}
	val, err := s.Pop()
	if err != nil {
		return err
	}
	vals := append(quot.Quotation, val)
	s.Push(data.NewQuotation(vals...))
	return nil
}



// Dictionary manipulation.


// let pops a string and the next value, then saves a word named for that string
// which evaluates to that value.
func let(d *data.Dictionary, s *data.Stack) error {
	wordName, err := s.PopType(data.String)
	if err != nil {
		return err
	}
	definition, err := s.Pop()
	if err != nil {
		return err
	}

	d.Set(wordName.Str, definition)
	return nil
}


// define pops a string and the next value, then saves a word named for that string which evaluates to a procedure that applies that value.
func define(d *data.Dictionary, s *data.Stack) error {

	wordName, err := s.PopType(data.String)
	if err != nil {
		return err
	}
	definition, err := s.Pop()
	if err != nil {
		return err
	}
	proc := data.NewProc(
		func(d *data.Dictionary, s *data.Stack) error {
			s.Push(definition)
			return apply(d, s)
		},
	)
	d.Set(wordName.Str, proc)

	return nil
}



// Stack manipulation.


// dup pops a value, then pushes it twice.
func dup(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}
	s.Push(val)
	s.Push(val)
	return nil
}


// drop discards the top value on the stack
func drop(d *data.Dictionary, s *data.Stack) error {
	_, err := s.Pop()
	if err != nil {
		return err
	}
	return nil
}


// lambda pops a value and pushes a procedure that evaluates to that value.
func lambda(d *data.Dictionary, s *data.Stack) error {
	head, err := s.Pop()
	if err != nil {
		return err
	}
	proc := data.NewProc(
		func(d *data.Dictionary, s *data.Stack) error {
			s.Push(head)
			return apply(d, s)
		},
	)
	s.Push(proc)
	return nil
}
