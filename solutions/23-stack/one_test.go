package stack

import "testing"

func TestCreateDefault(t *testing.T) {
	if s := Create(); cap(s.items) != 10 {
		t.Log("Create() did not create default stack of correct size")
		t.Fail()
	}
}

func TestCreate(t *testing.T) {
	t.Log("Setup!") // common setup code can go here
	t.Run("create", func(t *testing.T) {
		size := 25
		if s := Create(size); cap(s.items) != size {
			t.Log("Create(...) did not create stack of correct size")
			t.Fail()
		}
	})
	t.Run("create", func(t *testing.T) {
		size := 25
		if s := Create(size); cap(s.items) != size {
			t.Log("Create() did not create stack of correct size")
			t.Fail()
		}
	})
}
