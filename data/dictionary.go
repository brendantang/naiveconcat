package data

import ()

// A Dictionary binds words to data. When a word is evaluated, it is substituted
// for the data it binds in the Dictionary.
type Dictionary struct {
	bindings map[string]Value
}

// NewDictionary takes a map of word bindings and returns a Dictionary.
func NewDictionary(bindings map[string]Value) Dictionary {
	return Dictionary{bindings}
}

// Get retrieves the Value of a word.
func (dict Dictionary) Get(word string) (d Value, ok bool) {
	d, ok = dict.bindings[word]
	return
}

// Set saves the definition of a word.
func (dict Dictionary) Set(word string, val Value) (d Dictionary) {
	dict.bindings[word] = val
	return NewDictionary(dict.bindings)
}
