// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

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
