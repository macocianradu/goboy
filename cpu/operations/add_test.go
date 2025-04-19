package operations

import (
	"radu.macocian.me/goboy/memory"
	"testing"
)

func TestADD(t *testing.T) {
	r1 := byte(3)
	r2 := byte(4)
	ADD(&r1, &r2)

	expected := byte(7)
	actual := r1

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}

func TestADDFromMem(t *testing.T) {
	r1 := byte(3)
	addr := uint(1000)
	memory.Write8(addr, byte(6))
	ADDFromMem(&r1, addr)

	expected := byte(9)
	actual := r1

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}
