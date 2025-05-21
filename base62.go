package base62

import (
	"fmt"
	"math/big"
)

func Encode(s string, base int) (string, error) {
	var b big.Int
	if _, ok := b.SetString(s, base); !ok {
		return "", fmt.Errorf("invalid base%d string", base)
	}
	return b.Text(62), nil
}

func Decode(s string, base int) (string, error) {
	var b big.Int
	if _, ok := b.SetString(s, 62); !ok {
		return "", fmt.Errorf("invalid base62 string: %s", s)
	}
	return b.Text(base), nil
}
