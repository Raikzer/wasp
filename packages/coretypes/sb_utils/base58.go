package sb_utils

import "github.com/mr-tron/base58"

type base58Util struct{}

func (u base58Util) Decode(s string) ([]byte, error) {
	return base58.Decode(s)
}

func (u base58Util) Encode(data []byte) string {
	return base58.Encode(data)
}
