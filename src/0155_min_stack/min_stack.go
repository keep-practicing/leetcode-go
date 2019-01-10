/*
155. Min Stack
https://leetcode.com/problems/min-stack/

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
*/
// time: 2019-01-10

package minstack

// MinStack a stack that supports push, pop, top, and retrieving the minimum element in constant time.
type MinStack struct {
	data []int
}

// Constructor initialize your data structure here.
func Constructor() MinStack {
	return MinStack{}
}

// Push element x onto stack.
func (s *MinStack) Push(x int) {
	s.data = append(s.data, x)
}

// Pop Removes the element on top of the stack.
func (s *MinStack) Pop() {
	if !s.IsEmpty() {
		s.data = s.data[:s.GetSize()-1]
	}
}

// Top  Get the top element.
func (s *MinStack) Top() int {
	//if s.IsEmpty(){
	//	return 0
	//}
	// stack must have element.
	return s.data[s.GetSize()-1]
}

// GetMin retrieving the minimum element in constant time.
// time complexity: O(n)
func (s *MinStack) GetMin() int {
	//if s.IsEmpty() {
	//	return 0
	//}
	// stack must have element.
	stackSize := s.GetSize()
	ret := s.data[stackSize-1]

	for i := stackSize - 2; i >= 0; i-- {
		if s.data[i] < ret {
			ret = s.data[i]
		}
	}
	return ret
}

// GetSize get size of the stack.
func (s MinStack) GetSize() int {
	return len(s.data)
}

func (s MinStack) IsEmpty() bool {
	return s.GetSize() == 0
}
