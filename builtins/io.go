package builtins

import (
	"fmt"
	"github.com/brendantang/naiveconcat/data"
)

func say(d *data.Dictionary, s *data.Stack) error {
	val, err := s.Pop()
	if err != nil {
		return err
	}
	fmt.Println(val)
	return nil
}
