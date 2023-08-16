package flag

import (
	"reflect"
	"testing"
)

func TestStringLAngleSeparatedMapStringString(t *testing.T) {
	var nilMap map[string]string
	cases := []struct {
		desc   string
		m      *LAngleSeparatedMapStringString
		expect string
	}{
		{
			desc:   "nil",
			m:      NewLAngleSeparatedMapStringString(&nilMap),
			expect: "",
		},
		{
			desc:   "empty",
			m:      NewLAngleSeparatedMapStringString(&map[string]string{}),
			expect: "",
		},
		{
			desc:   "one key",
			m:      NewLAngleSeparatedMapStringString(&map[string]string{"one": "foo"}),
			expect: "one<foo",
		},
		{
			desc: "two keys",
			m: NewLAngleSeparatedMapStringString(&map[string]string{
				"one": "foo",
				"two": "bar",
			}),
			expect: "one<foo,two<bar",
		},
	}
	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := c.m.String()
			if c.expect != got {
				t.Fatalf("expect %q but got %q", c.expect, got)
			}
		})
	}
}

func TestSetLAngleSeparatedMapStringString(t *testing.T) {
	var nilMap map[string]string
	cases := []struct {
		desc   string
		vals   []string
		start  *LAngleSeparatedMapStringString
		expect *LAngleSeparatedMapStringString
		err    string
	}{
		// we initialize the map with a default key that should be cleared by Set
		{
			"clear defaults",
			[]string{""},
			NewLAngleSeparatedMapStringString(&map[string]string{"default": ""}),
			&LAngleSeparatedMapStringString{
				Map:         &map[string]string{},
				initialized: true,
			},
			"",
		},
		// make sure we still allocate for "initialized" maps where Map was initially set to a nil map
		{
			"allocates map if currently nil",
			[]string{""},
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &nilMap},
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{},
			},
			"",
		},
		// for most cases, we just reuse nilMap, which should be allocated by Set, and is reset before each test case
		{
			"empty",
			[]string{""},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{},
			},
			"",
		},
		{
			"one key",
			[]string{"one<foo"},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo"},
			},
			"",
		},
		{
			"two keys",
			[]string{"one<foo,two<bar"},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
			},
			"",
		},
		{
			"two keys, multiple Set invocations",
			[]string{"one<foo", "two<bar"},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
			},
			"",
		},
		{
			"two keys with space",
			[]string{"one<foo, two<bar"},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{"one": "foo", "two": "bar"},
			},
			"",
		},
		{
			"empty key",
			[]string{"<foo"},
			NewLAngleSeparatedMapStringString(&nilMap),
			&LAngleSeparatedMapStringString{
				initialized: true,
				Map:         &map[string]string{"": "foo"},
			},
			"",
		},
		{
			"missing value",
			[]string{"one"},
			NewLAngleSeparatedMapStringString(&nilMap),
			nil,
			"malformed pair, expect string<string",
		},
		{
			"no target",
			[]string{"a:foo"},
			NewLAngleSeparatedMapStringString(nil),
			nil,
			"no target (nil pointer to map[string]string)",
		},
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
