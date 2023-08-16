package flag

import (
	"fmt"
	"sort"
	"strings"
)

// ColonSeparatedMultiMapStringString supports setting a map[string][]string from an encoding
// that separates keys from values with ':' and separates key-value paris with ','.
// A key can be repeated multiple times, in which case the values are appended to a
// slice of strings associated with that key. Items in the list associated with a given
// key will appear in the order provided.
// For example: `a:hello,b:again,c:world,b:beautiful` results in `{"a": ["hello"], "b": ["again", "beautiful"], "c":
// ["world"]}`
// The first call to Set will clear the map before adding entries; subsequent calls will simply append to the map.
// This makes it possible to override default values with a command-line option rather than appending to defaults,
// while still allowing the distribution of key-value paris across multiple flag invocations.
// For example: `--flag "a:hello" --flag "b:again" --flag "b:beautiful" --flag "c:world"` results in `{"a": ["hello"],
// "b": ["again", "beautiful"], "c": ["world"]}`.
type ColonSeparatedMultiMapStringString struct {
	MultiMap    *map[string][]string
	initialized bool // set to true after the first Set call
}

// NewColonSeparatedMultiMapStringString takes a pointer to a map[string][]string and returns the
// ColonSeparatedMultiMapStringString flag parsing shim for that map.
func NewColonSeparatedMultiMapStringString(m *map[string][]string) *ColonSeparatedMultiMapStringString {
	return &ColonSeparatedMultiMapStringString{MultiMap: m}
}

// Set implements github.com/spf13/pflag.Value.
func (m *ColonSeparatedMultiMapStringString) Set(value string) error {
	if m.MultiMap == nil {
		return fmt.Errorf("no target (nil pointer to map[string][]string)")
	}
	if !m.initialized || *m.MultiMap == nil {
		// clear default values, or allocate if no existing map
		*m.MultiMap = make(map[string][]string)
		m.initialized = true
	}
	for _, pair := range strings.Split(value, ",") {
		if len(pair) == 0 {
			continue
		}
		kv := strings.SplitN(pair, ":", 2)
		if len(kv) != 2 {
			return fmt.Errorf("malformed pair, expect string:string")
		}
		(*m.MultiMap)[strings.TrimSpace(kv[0])] = append((*m.MultiMap)[strings.TrimSpace(kv[0])], strings.TrimSpace(kv[1]))
	}

	return nil
}

// String implements github.com/spf13/pflag.Value.
func (m *ColonSeparatedMultiMapStringString) String() string {
	type kv struct {
		k string
		v string
	}
	kvs := make([]kv, 0, len(*m.MultiMap))
	for k, vs := range *m.MultiMap {
		for i := range vs {
			kvs = append(kvs, kv{k: k, v: vs[i]})
		}
	}

	// stable sort by keys, order of values should be preserved
	sort.SliceStable(kvs, func(i, j int) bool {
		return kvs[i].k < kvs[j].k
	})

	pairs := make([]string, 0, len(kvs))
	for i := range kvs {
		pairs = append(pairs, fmt.Sprintf("%s:%s", kvs[i].k, kvs[i].v))
	}

	return strings.Join(pairs, ",")
}

// Type implements github.com/spf13/pflag.Value.
func (m *ColonSeparatedMultiMapStringString) Type() string {
	return "colonSeparatedMultiMapStringString"
}

// Empty implements OmitEmpty.
func (m *ColonSeparatedMultiMapStringString) Empty() bool {
	return len(*m.MultiMap) == 0
}
