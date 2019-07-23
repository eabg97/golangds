package golangds

import "fmt"

// Stack data structure
type Stack struct {
	storage
}

// Top func
func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	if s.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return *s.store[len(s.store)-1], nil
}

// Bottom func
func (s *Stack) Bottom() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	if s.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return *s.store[0], nil
}

// Push func
func (s *Stack) Push(element interface{}) {
	if s.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = append(s.store, &element)
}

// Pop func
func (s *Stack) Pop() error {
	if s.Empty() {
		return fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	if s.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = s.store[:len(s.store)-1]
	return nil
}

// Swap func
func (s *Stack) Swap(otherStack *Stack) {
	temp := otherStack
	otherStack = s
	s = temp
}
