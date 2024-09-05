# GO 1.23: implementation of go v1.23 for range feature


```go
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
```

## RUN it

```go
func main() {
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

	fmt.Println()

	for i, j := range s.All2() {
		fmt.Println(i, j)
		if i > 1 {
			break
		}
	}
}
```