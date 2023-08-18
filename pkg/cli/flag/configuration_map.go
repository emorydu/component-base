// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package flag

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

type ConfigurationMap map[string]string

var _ flag.Value = &ConfigurationMap{}

func (m *ConfigurationMap) String() string {
	var paris []string
	for k, v := range *m {
		paris = append(paris, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(paris)

	return strings.Join(paris, ",")
}

func (m *ConfigurationMap) Set(value string) error {
	for _, s := range strings.Split(value, ",") {
		if len(s) == 0 {
			continue
		}
		arr := strings.SplitN(s, "=", 2)
		if len(arr) == 2 {
			(*m)[strings.TrimSpace(arr[0])] = strings.TrimSpace(arr[1])
		} else {
			(*m)[strings.TrimSpace(arr[0])] = ""
		}
	}

	return nil
}

func (*ConfigurationMap) Type() string {
	return "mapStringString"
}
