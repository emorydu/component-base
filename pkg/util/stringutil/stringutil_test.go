package stringutil

import "testing"

func TestDiff(t *testing.T) {
	tests := [][]string{
		{"foo", "bar", "hello"},
		{"foo", "bar", "world"},
	}

	result := Diff(tests[0], tests[1])
	if len(result) != 1 || result[0] != "hello" {
		t.Fatalf("Diff failed")
	}
}
