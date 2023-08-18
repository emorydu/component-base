// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stringutil

import (
	"github.com/asaskevich/govalidator"
	"unicode/utf8"
)

// Diff creates a slice of slice values not included in the other given slice.
func Diff(base, exclude []string) []string {
	var result []string

	excludeMap := make(map[string]bool)
	for _, s := range exclude {
		excludeMap[s] = true
	}
	for _, s := range base {
		if !excludeMap[s] {
			result = append(result, s)
		}
	}

	return result
}

func Unique(ss []string) []string {
	result := make([]string, 0, len(ss))

	tmp := map[string]struct{}{}
	for _, s := range ss {
		if _, ok := tmp[s]; !ok {
			tmp[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}

func CamelCaseToUnderscore(str string) string {
	return govalidator.CamelCaseToUnderscore(str)
}

func UnderscoreToCamelCase(str string) string {
	return govalidator.UnderscoreToCamelCase(str)
}

func StringIn(str string, vals []string) bool {
	return FindString(vals, str) > -1
}

func FindString(vals []string, s string) int {
	for i, v := range vals {
		if s == v {
			return i
		}
	}

	return -1
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)

	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}

	return string(buf)
}
