package data

import (
	"fmt"
)

// A Dictionary binds words to data. When a word is evaluated, it is substituted
// for the data it binds in the Dictionary.
type Dictionary struct {
	parent   *Dictionary
	bindings map[string]Value
}

// NewDictionary takes a map of word bindings and returns a Dictionary.
func NewDictionary(parent *Dictionary, bindings map[string]Value) *Dictionary {
	return &Dictionary{parent, bindings}
}

// Get retrieves the Value of a word. If a definition isn't found, look
// recursively in each parent Dictionary.
func (dict *Dictionary) Get(word string) (d Value, ok bool) {
	d, ok = dict.bindings[word]
	if !ok && dict.parent != nil {
		return dict.parent.Get(word)
	}
	return
}

// Set saves the definition of a word.
func (dict *Dictionary) Set(word string, val Value) {
	dict.bindings[word] = val
}

func (dict *Dictionary) String() (s string) {
	for w, def := range dict.bindings {
		s = fmt.Sprintf("%s\n%s:\t%s", s, w, def)
	}
	return
}
