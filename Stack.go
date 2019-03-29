package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Stack struct {
	value []interface{}
	size  int
}

func (s *Stack) push(e int) {
	s.value = append(s.value, e)
	s.size += 1
}

func (s *Stack) init(len int) {
	s.value = make([]interface{}, len)
}

func (s *Stack) pop() (int, error) {
	ret := s.value[s.size-1]
	s.value = append(s.value[:s.size-1], s.value[s.size:]...)
	s.size = s.size - 1
	v, ok := ret.(int)

	if ok {
		return v, nil
	}
	return 0, errors.New("")
}

func main() {
	s := Stack{}
	length := 0
	s.init(length)
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.push(6)
	fmt.Println(s.value)
	fmt.Println(s.size)
	fmt.Println(s.pop())
	fmt.Println(s.value)
	s.push(7)
	fmt.Println(s.value)

	g := make([]interface{}, 0)
	g = append(g, 1)
	g = append(g, "a")
	g = append(g, s)

	for _, v := range g {
		value, ok := v.(int)
		switch v.(type) {
		case int:
			fmt.Printf("int type:%v\n", v)
		case string:
			fmt.Printf("string type:%v\n", v)
		case Stack:
			fmt.Printf("%v", v)
		}

		if ok {
			fmt.Println(value)
		}
	}
}
