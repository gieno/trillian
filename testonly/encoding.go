package testonly

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/google/trillian"
)

// MustDecodeBase64 expects a base 64 encoded string input and panics if it cannot be decoded
func MustDecodeBase64(b64 string) trillian.Hash {
	r, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(r)
	}
	return r
}

// MustHexDecode decodes its input string from hex and panics if this fails
func MustHexDecode(b string) trillian.Hash {
	r, err := hex.DecodeString(b)
	if err != nil {
		panic(err)
	}
	return r
}
