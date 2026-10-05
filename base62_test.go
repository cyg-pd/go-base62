package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase62(t *testing.T) {
	s, err := Encode("9e3635adb80a4e138220de9f6d12d1cb", 16)

	assert.NoError(t, err)
	assert.Equal(t, "4OxyOCWOdblodfDYb7qMyf", s)
}

func TestFromBase62(t *testing.T) {
	s, err := Decode("4OxyOCWOdblodfDYb7qMyf", 16)

	assert.NoError(t, err)
	assert.Equal(t, "9e3635adb80a4e138220de9f6d12d1cb", s)
}

func TestEncodeInvalidBase(t *testing.T) {
	_, err := Encode("ff", 1)
	assert.Error(t, err)
	_, err = Encode("ff", 63)
	assert.Error(t, err)
}

func TestDecodeInvalidBase(t *testing.T) {
	_, err := Decode("1A", 1)
	assert.Error(t, err)
	_, err = Decode("1A", 100)
	assert.Error(t, err)
}
