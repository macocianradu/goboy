package operations

import (
	"testing"
)

func TestShiftLeft(t *testing.T) {
	r1 := byte(0b01010101)
	ShiftLeft(&r1)

	expected := byte(0b10101010)
	actual := r1

	if actual != expected {
		t.Errorf("actual %08b != expected %08b", actual, expected)
	}
}

func TestShiftRight(t *testing.T) {
	r1 := byte(0b01010101)
	ShiftRight(&r1)

	expected := byte(0b00101010)
	actual := r1

	if actual != expected {
		t.Errorf("actual %08b != expected %08b", actual, expected)
	}
}
