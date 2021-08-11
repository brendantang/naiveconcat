package eval

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
	"sort"
)

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
