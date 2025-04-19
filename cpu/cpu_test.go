package cpu

import (
	"testing"
)

func TestCPUAF(t *testing.T) {
	cpu := CPU_struct{}
	cpu.A = 0x0A
	cpu.F = 0x0B

	expected := uint16(0x0A0B)
	actual := cpu.AF()

	if actual != expected {
		t.Errorf("actual %x != expected %x", actual, expected)
	}
}

func TestCPU_struct_SetAF(t *testing.T) {
	cpu := CPU_struct{}
	cpu.SetAF(0xABCD)

	expectedA := uint8(0xAB)
	expectedF := uint8(0xCD)
	actualA := cpu.A
	actualF := cpu.F
	if actualA != expectedA {
		t.Errorf("actual %x != expected %x", actualA, expectedA)
	}
	if actualF != expectedF {
		t.Errorf("actual %x != expected %x", actualF, expectedF)
	}
}

func TestCPU_struct_GetAndSetAF(t *testing.T) {
	cpu := CPU_struct{}
	cpu.SetAF(0xABCD)

	expectedAF := uint16(0xABCD)
	actualAF := cpu.AF()

	if actualAF != expectedAF {
		t.Errorf("actual %x != expected %x", actualAF, expectedAF)
	}
}
