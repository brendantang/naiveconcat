package builtins

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

func say(dict data.Dictionary, s *data.Stack) (data.Dictionary, error) {
	d, err := s.Pop()
	if err != nil {
		return dict, err
	}
	fmt.Println(d)
	return dict, nil
}
