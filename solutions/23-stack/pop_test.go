package stack

import "testing"

func TestPushPop(t *testing.T) {
	s := Stack{}
	s.Push(1)
	if val, ok := s.Pop(); !ok || val != 1 {
		t.Log("Pop() failed to return 1")
		t.Fail()
	}
}
