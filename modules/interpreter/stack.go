package interpreter

import "sync"

// Stack struct which contains a list of Items
type Stack struct {
	items []interface{}
	mutex sync.Mutex
}

// NewEmptyStack() returns a new instance of Stack with zero elements
func NewEmptyStack() *Stack {
	return &Stack{
		items: nil,
	}
}

// NewStack() returns a new instance of Stack with list of specified elements
func NewStack(items []interface{}) *Stack {
	return &Stack{
		items: items,
	}
}

// Push() adds new item to top of existing/empty stack
func (stack *Stack) Push(item interface{}) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, item)
}

// Pop() removes most recent item(top) from stack
func (stack *Stack) Pop() interface{} {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if len(stack.items) == 0 {
		return -1
	}
	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return lastItem
}

// IsEmpty() returns whether the stack is empty or not (boolean result)
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items) == 0
}

// Top() returns the last inserted item in stack without removing it.
func (stack *Stack) Top() interface{} {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if len(stack.items) == 0 {
		return -1
	}
	return stack.items[len(stack.items)-1]
}

// Size() returns the number of items currently in the stack
func (stack *Stack) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	return len(stack.items)
}

// Clear() removes all items from the stack
func (stack *Stack) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = nil
}

// Contains() returns whether the stack contains the specified item or not (boolean result)
func (stack *Stack) Contains(item interface{}) bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	for _, stackItem := range stack.items {
		if stackItem == item {
			return true
		}
	}
	return false
}
