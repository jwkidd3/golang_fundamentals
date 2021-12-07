package testing

import (
	"strings"
	"testing"
)
func Test_Index(t *testing.T) {
	t.Parallel()
	sentence := "Gophers are amazing"
	substring := "are"
	got := strings.Index(sentence, substring)
	// expected value:
	exp := 8
	if got != exp {
		t.Errorf("unexpected value for indexing %q for %q.  got %d, exp %d", sentence, substring, got, exp)
	}
}

func Test_First(t *testing.T) {
	t.Parallel()
	sentence := "Gophers are amazing"
	substring := "Gophers"
	got := strings.Index(sentence, substring)
	// expected value:
	exp := 0
	if got != exp {
		t.Errorf("unexpected value for indexing %q for %q.  got %d, exp %d", sentence, substring, got, exp)
	}
}

func Test_None(t *testing.T) {
	t.Parallel()
	sentence := "Gophers are amazing"
	substring := "rust"
	got := strings.Index(sentence, substring)
	// expected value:
	exp := -1
	if got != exp {
		t.Errorf("unexpected value for indexing %q for %q.  got %d, exp %d", sentence, substring, got, exp)
	}
}