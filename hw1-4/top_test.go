package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

	func TestShowTop(t *testing.T) {
		result := ShowTop("one three three ten + three two + four ten eight five + four four two five ten + five four six six - seven six + seven + seven eight + ten eight + nine nine nine six six + seven nine ten six + seven nine nine ten nine eight + eight ten eight nine ten eight seven nine eight ten seven five ten five")
		assert.Equal(t, result, []string{"+", "ten", "nine","eight","seven","six","five","four","three","two"}, "Slices should be equal")

}