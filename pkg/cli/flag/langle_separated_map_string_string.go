package flag

import (
	"fmt"
	"sort"
	"strings"
)

// LAngleSeparatedMapStringString can be set from the command line with the format `--flag "string<string"`.
// Multiple comma-separated key-value paris in a single invocation are supported. For example: `--flag "a<foo,b<bar"`.
// Multiple flag invocations are supported. For example: `--flag "a<foo" --flag "b<foo"`.
type LAngleSeparatedMapStringString struct {
	Map         *map[string]string
	initialized bool // set to true first Set call
}

// NewLAngleSeparatedMapStringString takes a pointer to a map[string]string and returns the
// LAngleSeparatedMapStringString flag parsing shim for that map.
func NewLAngleSeparatedMapStringString(m *map[string]string) *LAngleSeparatedMapStringString {
	return &LAngleSeparatedMapStringString{Map: m}
}

// String implements github.com/spf13/pflag.Value.
func (m *LAngleSeparatedMapStringString) String() string {
	var paris []string
	for k, v := range *m.Map {
		paris = append(paris, fmt.Sprintf("%s<%s", k, v))
	}
	sort.Strings(paris)

	return strings.Join(paris, ",")
}

// Set implements github.com/spf13/pflag.Value.
func (m *LAngleSeparatedMapStringString) Set(value string) error {
	if m.Map == nil {
		return fmt.Errorf("no target (nil pointer to map[string]string)")
	}
	if !m.initialized || *m.Map == nil {
		// clear default values, or allocate if no existing map
		*m.Map = make(map[string]string)
		m.initialized = true
	}
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "<", 2)
		if len(arr) != 2 {
			return fmt.Errorf("malformed pair, expect string<string")
		}
		(*m.Map)[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
	}

	return nil
}

// Type implements github.com/spf13/pflag.Value.
func (*LAngleSeparatedMapStringString) Type() string {
	return "mapStringString"
}

// Empty implements OmitEmpty.
func (m *LAngleSeparatedMapStringString) Empty() bool {
	return len(*m.Map) == 0
}
