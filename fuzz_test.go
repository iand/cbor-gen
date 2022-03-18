//go:build go1.18
// +build go1.18

package typegen_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/whyrusleeping/cbor-gen"
)

func fromHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func FuzzCborCidUnmarshalMarshal(f *testing.F) {
	for _, b := range [][]byte{
		// Cbor encoding of a valid cid
		fromHex("d82a582500015512209d8453505bdc6f269678e16b3e56c2a2948a41f2c792617cc9611ed363c95b63"),
	} {
		f.Add(b)
	}
	f.Fuzz(func(t *testing.T, data []byte) {
		var c typegen.CborCid
		if err := c.UnmarshalCBOR(bytes.NewReader(data)); err != nil {
			return // invalid cid bytes, do not re-encode
		}

		var buf bytes.Buffer
		if err := c.MarshalCBOR(&buf); err != nil {
			t.Fatalf("re-encode of valid cid failed: %v", err)
		}
	})
}
