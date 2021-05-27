package stack

import "testing"

func TestPush(t *testing.T) {
	s := Stack{}
	item := 1
	for i := 1; i <= MaxSize; i++ {
		s.Push(item)
	}
	if ok := s.Push(item); ok {
		t.Log("Push() allowed when stack was full")
		t.Fail()
	}
	if val, ok := s.Top(); !ok || val != item {
		t.Log("Pop() failed to return pushed item")
		t.Fail()
	}
}
