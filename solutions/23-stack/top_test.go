package stack

import "testing"

func TestTop(t *testing.T) {
	s := Stack{} // cf. s := new(Stack)
	s.Push(1)
	if val, ok := s.Top(); ok && val != 1 {
		t.Log("Top() failed to return 1")
		t.Fail()
	}
}
