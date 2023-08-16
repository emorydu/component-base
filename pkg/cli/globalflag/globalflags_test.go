package globalflag

import (
	"flag"
	cliflag "github.com/emorydu/component-base/pkg/cli/flag"
	"github.com/spf13/pflag"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestAddGlobalFlags(t *testing.T) {
	namedFlagSets := cliflag.NamedFlagSets{}
	nfs := namedFlagSets.FlagSet("global")
	AddGlobalFlags(nfs, "test-cmd")

	var actualFlag []string
	nfs.VisitAll(func(flag *pflag.Flag) {
		actualFlag = append(actualFlag, flag.Name)
	})

	// Get all flags from flags.CommandLine, except flag `test.*`.
	wantedFlag := []string{"help"}
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.VisitAll(func(flag *pflag.Flag) {
		if !strings.Contains(flag.Name, "test.") {
			wantedFlag = append(wantedFlag, normalize(flag.Name))
		}
	})
	sort.Strings(wantedFlag)

	if !reflect.DeepEqual(wantedFlag, actualFlag) {
		t.Errorf("[Default]: expected %+v, got %+v", wantedFlag, actualFlag)
	}

	tests := []struct {
		expectedFlag  []string
		matchExpected bool
	}{

		{
			// Happy case
			expectedFlag:  []string{"help"},
			matchExpected: false,
		},
		{
			// Missing flag
			expectedFlag:  []string{"logtostderr", "log-dir"},
			matchExpected: true,
		},
		{
			// Empty flag
			expectedFlag:  []string{},
			matchExpected: true,
		},
		{
			// Invalid flag
			expectedFlag:  []string{"foo"},
			matchExpected: true,
		},
	}

	for i, tt := range tests {
		if reflect.DeepEqual(tt.expectedFlag, actualFlag) == tt.matchExpected {
			t.Errorf("[%d]: expected %+v, got %+v", i, tt.expectedFlag, actualFlag)
		}
	}
}
