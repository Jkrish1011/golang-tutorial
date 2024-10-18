package leetcode

type Comparable interface {
	~int | ~string | ~float64 | ~struct{} | any
}

type Stack[T Comparable] []T

func NewStack[T Comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	topIdx := len(*s) - 1
	retVal := (*s)[topIdx]
	*s = (*s)[:topIdx]
	return retVal
}

func (s *Stack[T]) IsEmpty() bool {
	return len((*s)) == 0
}
