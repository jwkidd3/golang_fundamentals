// A simple stack object with 10 slots for integers
package stack

const MaxSize = 10

// Stack object
type Stack struct {
	Ptr   int
	Items [MaxSize]int
}

// Push an item onto the stack. Return false if already full.
func (s *Stack) Push(i int) bool {
	if s.Size() == MaxSize {
		return false
	}
	s.Items[s.Ptr] = i
	s.Ptr++
	return true
}

// Pop an item from the stack, return false if empty
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	s.Ptr--
	return s.Items[s.Ptr], true
}

// Return the number of items in the stack
func (s *Stack) Size() int {
	return s.Ptr
}

// Return whether or not the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

// Return top item from stack along with boolean indication of success.
// If stack empty, return false.
func (s *Stack) Top() (int, bool) {
	if !s.IsEmpty() {
		return s.Items[s.Ptr-1], true
	}
	return 0, false
}
