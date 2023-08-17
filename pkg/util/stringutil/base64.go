package stringutil

import (
	"bytes"
	"encoding/base64"
	"io"
)

func DecodeBase64(s string) ([]byte, error) {
	return io.ReadAll(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(s)))
}
