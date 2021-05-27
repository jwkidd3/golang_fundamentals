package stack

import "testing"

func TestEmpty(t *testing.T) {
	s := Stack{} // cf. s := new(Stack)
	if !s.IsEmpty() {
		t.Log("IsEmpty() failed to return true (1)")
		t.Fail()
	}
	s.Push(1)
	if _, ok := s.Pop(); ok {
		if !s.IsEmpty() {
			t.Log("IsEmpty() failed to return true (2)")
			t.Fail()
		}
	} else {
		t.Log("Pop() failed in IsEmpty() test")
		t.Fail()
	}
}
