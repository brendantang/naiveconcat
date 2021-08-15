package data

import (
	"fmt"
	"strings"
)

// A Dictionary binds words to data. When a word is evaluated, it is substituted
// for the data it binds in the Dictionary.
type Dictionary struct {
	parent   *Dictionary
	Bindings map[string]Value
}

func (dict *Dictionary) String() (s string) {
	for w, def := range dict.Bindings {
		s = fmt.Sprintf("%s\n%s:\t%s", s, w, def)
	}
	return
}

// NewDictionary takes a map of word bindings and returns a Dictionary.
func NewDictionary(parent *Dictionary, bindings map[string]Value) *Dictionary {
	return &Dictionary{parent, bindings}
}

// Get retrieves the Value of a word. If a definition isn't found, look
// recursively in each parent Dictionary.
func (dict *Dictionary) Get(word string) (d Value, ok bool) {
	d, ok = dict.Bindings[word]
	if !ok && dict.parent != nil {
		return dict.parent.Get(word)
	}
	return
}

// Set saves the definition of a word. If the word name contains any white space, it
// is replaced with a single underscore.
func (dict *Dictionary) Set(word string, val Value) {
	word = strings.Join(strings.Fields(word), "_")
	dict.Bindings[word] = val
}

// Depth returns the number of parent Dictionaries wrapping dict.
func (dict *Dictionary) Depth() int {
	if dict.parent == nil {
		return 0
	}
	return 1 + dict.parent.Depth()
}
