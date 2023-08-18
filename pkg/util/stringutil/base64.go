// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package stringutil

import (
	"bytes"
	"encoding/base64"
	"io"
)

func DecodeBase64(s string) ([]byte, error) {
	return io.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(s)))
}
