package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	unpacked, _:= Unpack("a4bc2d5e")
	assert.Equal(t, "aaaabccddddde", unpacked, "Unpacked stings should be matched")
	unpacked, _ = Unpack("abcd")
	assert.Equal(t, "abcd", unpacked, "Unpacked stings should be matched")
	_, err := Unpack("44234234")
	assert.NotNil(t, err)
	_, err = Unpack("4423s4234")
	assert.Nil(t, err)


}