package operations

import (
	"testing"
)

func TestINC(t *testing.T) {
	r1 := byte(3)
	INC(&r1)

	expected := byte(4)
	actual := r1

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}

func TestINC16(t *testing.T) {
	r1 := byte(0xAB)
	r2 := byte(0x11)
	INC16(&r1, &r2)

	expected := uint16(0xAB12)
	actual := (uint16(r1) << 8) | uint16(r2)
	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}

	r1 = byte(0x11)
	r2 = byte(0xFF)
	INC16(&r1, &r2)

	expected = uint16(0x1200)
	actual = (uint16(r1) << 8) | uint16(r2)
	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}
