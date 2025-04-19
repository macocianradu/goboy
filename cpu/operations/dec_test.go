package operations

import "testing"

func TestDEC(t *testing.T) {
	r1 := byte(3)
	DEC(&r1)

	expected := byte(2)
	actual := r1

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}

func TestDEC16(t *testing.T) {
	r1 := byte(0xAB)
	r2 := byte(0x11)
	DEC16(&r1, &r2)

	expected := uint16(0xAB10)
	actual := (uint16(r1) << 8) | uint16(r2)
	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}

	r1 = byte(0x11)
	r2 = byte(0x00)
	DEC16(&r1, &r2)

	expected = uint16(0x1000)
	actual = (uint16(r1) << 8) | uint16(r2)
	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}
