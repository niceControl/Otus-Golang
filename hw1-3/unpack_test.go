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
	err, _ := Unpack("444")
	assert.Equal(t,"Incorrect input", err)

}