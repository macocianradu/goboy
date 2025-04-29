package operations

import (
	"testing"
)

func TestShift(t *testing.T) {
	r1 := byte(0b01010101)
	Shift(&r1)

	expected := byte(0b10101010)
	actual := r1

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}
