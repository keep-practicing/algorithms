package stack

import "errors"

// Stack 栈
type Stack []interface{}

// Len 栈大小
func (stack Stack) Len() int {
	return len(stack)
}

// IsEmpty 是否为空
func (stack Stack) IsEmpty() bool {
	return stack.Len() == 0
}

// Cap  栈容量
func (stack Stack) Cap() int {
	return cap(stack)
}

// Push 压栈
func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

// Top  栈顶元素
func (stack Stack) Top() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

// Pop 出栈
func (stack *Stack) Pop() (interface{}, error) {
	if stack.Len() == 0 {
		return nil, errors.New("Out of index, len is 0")
	}

	value := (*stack)[stack.Len()-1]
	*stack = (*stack)[:stack.Len()-1]
	return value, nil
}
