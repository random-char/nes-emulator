package register_test

import (
	"nes-emulator/pkg/ppu/register"
	"testing"
)

func TestControlReg(t *testing.T) {
	control := register.NewControlReg()
	checkReg(t, uint16(control.GetReg()), 0b00_000_000)

	control.SetReg(0b00_000_000)
	control.GetNamtableX()
}

func TestStatusReg(t *testing.T) {
	status := register.NewStatusReg()
	checkReg(t, status.GetReg(), 0b00_000_000)

	status.SetReg(0b00_000_000)
}

func TestLoopyReg(t *testing.T) {
	loopy := register.NewLoopyReg()
	checkReg(t, loopy.GetReg(), 0b0_000_000_000_000_000)

	loopy.SetNametableX(1)
	checkReg(t, loopy.GetReg(), 0b0_000_010_000_000_000)

	loopy.SetNametableY(1)
	checkReg(t, loopy.GetReg(), 0b0_000_110_000_000_000)

	if loopy.GetNametableX() != 1 {
		t.Errorf("wrong nametableX: expected 1, got %b", loopy.GetNametableX())
	}

	loopy.SetNametableX(0)
	if loopy.GetNametableX() != 0 {
		t.Errorf("wrong nametableX: expected 0, got %b", loopy.GetNametableX())
	}
}

func checkReg[T uint8|uint16](
	t *testing.T,
	reg, expected T,
) {
	if reg != expected {
		t.Errorf(
			"wrong reg data: expected %16b, got %16b",
			expected,
			reg,
		)
	}
}
