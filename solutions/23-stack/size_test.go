package stack

import "testing"

func TestSize(t *testing.T) {
	s := Stack{}

	for i := 1; i <= MaxSize; i++ {
		s.Push(1)
	}
	if val, ok := s.Top(); !ok || val != 1 {
		t.Log("Pop() failed to return 1")
		t.Fail()
	}
	for i := 1; i < MaxSize; i++ {
		s.Push(1)
	}
	if ok := s.Push(1); ok {
		t.Log("Push() allowed when stack was full")
		t.Fail()
	}
}
