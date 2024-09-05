package main

import (
	"errors"
	"fmt"
	"iter"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type Stack[T any] struct {
	head *node[T]
}

func (s *Stack[T]) Push(v T) {
	s.head = &node[T]{value: v, next: s.head}
}

var ErrEmpty = errors.New("empty stack")

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var v T
		return v, ErrEmpty
	}

	n := s.head
	s.head = s.head.next
	return n.value, nil
}

func (s *Stack[T]) Do(yield func(v T)) {
	for n := s.head; n != nil; n = n.next {
		yield(n.value)
	}
}

func iter1(yeld func(i, j int) bool) {
	for i := range 6 {
		yeld(i, i+1)
	}
}

type Slice []int

func (s Slice) All() func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		for i := range s {
			if !yield(s[i]) {
				return
			}
		}
	}
}

func (s Slice) All2() func(yield func(i, j int) bool) {
	return func(yield func(i, j int) bool) {
		for i := range s {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func main() {
	// s := Stack[int]{}
	// s.Push(1)
	// s.Push(2)
	// s.Push(3)
	// s.Push(4)

	// s.Do(func(n int) {
	// 	fmt.Println(n)
	// })

	// for i := range 10 {
	// 	fmt.Println(i)
	// }
	// for k, v := range iter1 {
	// 	fmt.Println(k, v)
	// 	if k > 1 {
	// 		break
	// 	}
	// }

	s := Slice{1, 2, 3, 4, 5, 6}

	next, stop := iter.Pull2(s.All2())
	defer stop()

	for {
		k, v, valid := next()
		if !valid {
			break
		}
		fmt.Println(k, v)
	}

	// for i, j := range s.All2() {
	// 	fmt.Println(i, j)
	// 	if i > 1 {
	// 		break
	// 	}
	// }
	// return

}
