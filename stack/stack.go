package stack

import "fmt"

type Stack struct {
	vals []int
}

type StackEmptyError struct {
	message string
}

func (s *StackEmptyError) Error() string {
	return "Cant pop/peek on empty stack"
}

type IStack interface {
	pop() (int, error)

	push(val int)

	peek() (int, error)
}

func (s *Stack) pop() (int, error) {
	stackLength := len(s.vals)
	if stackLength == 0 {
		err := StackEmptyError{}
		return -1, &err
	} else {
		last := s.vals[stackLength-1]
		s.vals = s.vals[:stackLength-1]
		return last, nil
	}
}

func (s *Stack) push(val int) {
	s.vals = append(s.vals, val)
}

func (s *Stack) peek() (int, error) {
	stackLength := len(s.vals)
	if stackLength == 0 {
		return -1, &StackEmptyError{}
	} else {
		return s.vals[stackLength-1], nil
	}
}

func TestStack() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	popped, err := stack.pop()
	fmt.Println(popped, err)
	popped, err = stack.peek()
	fmt.Println(popped, err)
	popped, err = stack.pop()
	fmt.Println(popped, err)
	popped, err = stack.pop()
	fmt.Println(popped, err)
	popped, err = stack.pop()
	fmt.Println(popped, err)
	popped, err = stack.pop()
	fmt.Println(popped, err)
	popped, err = stack.peek()
	fmt.Println(popped, err)

}
