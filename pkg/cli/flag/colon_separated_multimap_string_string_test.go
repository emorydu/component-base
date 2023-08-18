// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package flag

import (
	"reflect"
	"testing"
)

func TestStringColonSeparatedMultiMapStringString(t *testing.T) {
	var nilMap map[string][]string
	cases := []struct {
		desc   string
		m      *ColonSeparatedMultiMapStringString
		expect string
	}{
		{"nil", NewColonSeparatedMultiMapStringString(&nilMap), ""},
		{"empty", NewColonSeparatedMultiMapStringString(&map[string][]string{}), ""},
		{"empty key", NewColonSeparatedMultiMapStringString(
			&map[string][]string{
				"": {"foo"},
			}), ":foo"},
		{"one key", NewColonSeparatedMultiMapStringString(
			&map[string][]string{
				"one": {"foo"},
			}), "one:foo"},
		{"two keys", NewColonSeparatedMultiMapStringString(
			&map[string][]string{
				"one": {"foo"},
				"two": {"bar"},
			},
		), "one:foo,two:bar"},
		{"two keys, multiple items in one key", NewColonSeparatedMultiMapStringString(
			&map[string][]string{
				"one": {"foo", "baz"},
				"two": {"bar"},
			}), "one:foo,one:baz,two:bar"},
		{"three keys, multiple items in one key", NewColonSeparatedMultiMapStringString(
			&map[string][]string{
				"a": {"hello"},
				"b": {"again", "beautiful"},
				"c": {"world"},
			}),
			"a:hello,b:again,b:beautiful,c:world"},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := c.m.String()
			if got != c.expect {
				t.Fatalf("expect %q but got %q", c.expect, got)
			}
		})
	}
}

func TestSetColonSeparatedMultiMapStringString(t *testing.T) {
	var nilMap map[string][]string
	cases := []struct {
		desc   string
		vals   []string
		start  *ColonSeparatedMultiMapStringString
		expect *ColonSeparatedMultiMapStringString
		err    string
	}{
		// we initialize the map with a default key that should be cleared by Set
		{
			desc:  "clears defaults",
			vals:  []string{""},
			start: NewColonSeparatedMultiMapStringString(&map[string][]string{"default": {}}),
			expect: &ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap:    &map[string][]string{},
			},
			err: "",
		},
		// make sure we still allocate for "initialized" multiMaps where MultiMap was initially set to a nil map
		{
			desc: "allocates map if currently nil",
			vals: []string{""},
			start: &ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap:    &nilMap,
			},
			expect: &ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap:    &map[string][]string{},
			},
			err: "",
		},
		// for most cases, we just reuse nilMap, which should be allocated by Set, and is reset before each test case
		{
			"empty",
			[]string{""},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap:    &map[string][]string{}}, ""},
		{
			"empty key",
			[]string{":foo"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"": {"foo"},
				}}, ""},
		{
			"one key",
			[]string{"one:foo"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"one": {"foo"},
				}}, ""},
		{
			"two keys",
			[]string{"one:foo,two:bar"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"one": {"foo"},
					"two": {"bar"},
				}}, ""},
		{
			"two keys with space",
			[]string{"one:foo, two:bar"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"one": {"foo"},
					"two": {"bar"},
				}}, ""},
		{
			"two keys, multiple items in one key",
			[]string{"one: foo, two:bar, one:baz"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"one": {"foo", "baz"},
					"two": {"bar"},
				}}, ""},
		{
			"three keys, multiple items in one key",
			[]string{"a:hello,b:again,c:world,b:beautiful"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"a": {"hello"},
					"b": {"again", "beautiful"},
					"c": {"world"},
				}}, ""},
		{
			"three keys, multiple items in one key, multiple Set invocations",
			[]string{"a:hello,b:again", "c:world", "b:beautiful"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			&ColonSeparatedMultiMapStringString{
				initialized: true,
				MultiMap: &map[string][]string{
					"a": {"hello"},
					"b": {"again", "beautiful"},
					"c": {"world"},
				}},
			"",
		},
		{
			"missing value",
			[]string{"a"},
			NewColonSeparatedMultiMapStringString(&nilMap),
			nil,
			"malformed pair, expect string:string"},
		{"no target", []string{"a:foo"},
			NewColonSeparatedMultiMapStringString(nil),
			nil,
			"no target (nil pointer to map[string][]string)"},
	}

	for _, c := range cases {
		nilMap = nil
		t.Run(c.desc, func(t *testing.T) {
			var err error
			for _, val := range c.vals {
				err = c.start.Set(val)
				if err != nil {
					break
				}
			}
			if c.err != "" {
				if err == nil || err.Error() != c.err {
					t.Fatalf("expect error %s but got %v", c.err, err)
				}
				return
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(c.expect, c.start) {
				t.Fatalf("expect %#v but got %#v", c.expect, c.start)
			}
		})
	}
}

func TestRoundTripColonSeparatedMultiMapStringString(t *testing.T) {
	cases := []struct {
		desc   string
		vals   []string
		expect string
	}{
		{"empty", []string{""}, ""},
		{"empty key", []string{":foo"}, ":foo"},
		{"one key", []string{"one:foo"}, "one:foo"},
		{"two keys", []string{"one:foo,two:bar"}, "one:foo,two:bar"},
		{"two keys, multiple items in one key", []string{"one:foo, two:bar, one:baz"}, "one:foo,one:baz,two:bar"},
		{
			"three keys, multiple items in one key",
			[]string{"a:hello,b:again,c:world,b:beautiful"},
			"a:hello,b:again,b:beautiful,c:world",
		},
		{
			"three keys, multiple items in one key, multiple Set invocations",
			[]string{"a:hello,b:again", "c:world", "b:beautiful"},
			"a:hello,b:again,b:beautiful,c:world",
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			m := NewColonSeparatedMultiMapStringString(&map[string][]string{})
			for _, val := range c.vals {
				if err := m.Set(val); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			got := m.String()
			if c.expect != got {
				t.Fatalf("expect %q but got %q", c.expect, got)
			}
		})
	}

}

func TestEmptyColonSeparatedMultiMapStringString(t *testing.T) {
	var nilMap map[string][]string
	cases := []struct {
		desc   string
		val    *ColonSeparatedMultiMapStringString
		expect bool
	}{
		{"nil", NewColonSeparatedMultiMapStringString(&nilMap), true},
		{"empty", NewColonSeparatedMultiMapStringString(&map[string][]string{}), true},
		{"populated", NewColonSeparatedMultiMapStringString(&map[string][]string{"foo": {}}), false},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := c.val.Empty()
			if got != c.expect {
				t.Fatalf("expect %t but got %t", c.expect, got)
			}
		})
	}
}
