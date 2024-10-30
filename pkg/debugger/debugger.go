package debugger

import (
	"fmt"
)

type Debugger struct {
	PrintData bool

	CpuPc          uint16
	CpuOpCode      uint8
	CpuOpName      string
	CpuNumOperands uint
	CpuOperand1    uint8
	CpuOperand2    uint8
	CpuRegA        uint8
	CpuRegX        uint8
	CpuRegY        uint8
	CpuStatus      uint8
	CpuSP          uint8
	CpuCycles      uint16
}

func (d *Debugger) Print() {
	var operands string
	switch d.CpuNumOperands {
	case 0:
		operands = "     "
		operands = "" //remove
	case 1:
		operands = fmt.Sprintf("%02X   ", d.CpuOperand1)
		operands = fmt.Sprintf("%02X", d.CpuOperand1) //remove
	case 2:
		operands = fmt.Sprintf("%02X %02X", d.CpuOperand1, d.CpuOperand2)
	}

	fmt.Printf(
		"%04X %02X %s %s A:%02X X:%02X Y:%02X P:%02X SP:%02X CYC:%d\n",
		d.CpuPc,
		d.CpuOpCode,
		operands,
		d.CpuOpName,
		d.CpuRegA,
		d.CpuRegX,
		d.CpuRegY,
		d.CpuStatus,
		d.CpuSP,
		d.CpuCycles,
	)
}

func (d *Debugger) Reset() {
	d.CpuPc = 0
	d.CpuOpCode = 0
	d.CpuOpName = ""
	d.CpuNumOperands = 0
	d.CpuOperand1 = 0
	d.CpuOperand2 = 0
	d.CpuRegA = 0
	d.CpuRegX = 0
	d.CpuRegY = 0
	d.CpuSP = 0
}
