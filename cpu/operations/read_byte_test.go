package operations

import (
	"testing"
)

func TestReadLowerByte(t *testing.T) {
	SP := uint16(0x1234)

	actual := GetLowerByte(SP)

	if actual != 0x34 {
		t.Errorf("actual %v != expected %v", actual, true)
	}
}

func TestReadUpperByte(t *testing.T) {
	SP := uint16(0x1234)

	actual := GetHigherByte(SP)

	if actual != 0x12 {
		t.Errorf("actual %v != expected %v", actual, true)
	}
}
