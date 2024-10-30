package nes

import (
	"bufio"
	"fmt"
	"nes-emulator/pkg/cartridge"
	"nes-emulator/pkg/cpu"
	"nes-emulator/pkg/debugger"
	"os"
	"strings"
	"testing"
)

func TestFromNestestRom(t *testing.T) {
	debugger := &debugger.Debugger{}
	cpu.InitialState.WithPc(0xC000)

	nes := New(
		nil,
		nil,
		debugger,
	)

	cart, err := cartridge.LoadFromFile("../../test/testdata/roms/nestest.nes")
	if err != nil {
		t.Fatal(err)
	}

	nestestLog, err := os.Open("../../test/testdata/logs/nestest.log")
	if err != nil {
		t.Fatal(err)
	}

	nes.InsertCartridge(cart)
	nes.Reset()

	//finish reset
	for !nes.cpu.OperationCompleted() {
		nes.Clock()
	}
	nes.Clock()
	nes.Clock()
	nes.Clock()

	scanner := bufio.NewScanner(nestestLog)

	for scanner.Scan() {
		logLine := scanner.Text()
		logData := strings.Fields(logLine)

		fmt.Println(logData)

		logIndex := 0
		pc := logData[logIndex]
		logIndex++

		opcode := logData[logIndex]
		logIndex++

		nes.Clock()
		nes.Clock()
		nes.Clock()
		for !nes.cpu.OperationCompleted() {
			nes.Clock()
		}
		debugger.Print()
		fmt.Println("-------------")

		debuggerPc := fmt.Sprintf("%04X", debugger.CpuPc)
		if debuggerPc != pc {
			fmt.Printf("%+v\n", debugger)
			t.Fatalf("Expected ps %s, got %s\n", pc, debuggerPc)
		}

		debuggerOpcode := fmt.Sprintf("%02X", debugger.CpuOpCode)
		if debuggerOpcode != opcode {
			fmt.Printf("%+v\n", debugger)
			t.Fatalf("Expected opcode %s, got %s\n", opcode, debuggerOpcode)
		}

		if debugger.CpuOpName == "NOP" || debugger.CpuOpName == "*NOP" {
			//failed num of operands assertions
			debugger.Reset()
			continue
		}

		var operand1, operand2 string = "", ""
		switch debugger.CpuNumOperands {
		case 0:
		case 1:
			operand1 = logData[logIndex]
			logIndex++

			debuggerOperand1 := fmt.Sprintf("%02X", debugger.CpuOperand1)
			if debuggerOperand1 != operand1 {
				t.Fatalf("Expected operand1 %s, got %s\n", operand1, debuggerOperand1)
			}
		case 2:
			operand1 = logData[logIndex]
			logIndex++

			debuggerOperand1 := fmt.Sprintf("%02X", debugger.CpuOperand1)
			if debuggerOperand1 != operand1 {
				t.Fatalf("Expected operand1 %s, got %s\n", operand1, debuggerOperand1)
			}

			operand2 = logData[logIndex]
			logIndex++

			debuggerOperand2 := fmt.Sprintf("%02X", debugger.CpuOperand2)
			if debuggerOperand2 != operand2 {
				t.Fatalf("Expected operand2 %s, got %s\n", operand2, debuggerOperand2)
			}
		default:
			t.Fatal("NumOperands should be 0, 1 or 2")
		}

		//opName := logData[logIndex]
		logIndex++

		a := logData[logIndex]
		logIndex++
		debugA := fmt.Sprintf("A:%02X", debugger.CpuRegA)
		if a != debugA {
			t.Fatalf("Expected a %s, got %s\n", a, debugA)
		}

		x := logData[logIndex]
		logIndex++
		debugX := fmt.Sprintf("X:%02X", debugger.CpuRegX)
		if x != debugX {
			t.Fatalf("Expected x %s, got %s\n", x, debugX)
		}

		y := logData[logIndex]
		logIndex++
		debugY := fmt.Sprintf("Y:%02X", debugger.CpuRegY)
		if y != debugY {
			t.Fatalf("Expected y %s, got %s\n", y, debugY)
		}

		p := logData[logIndex]
		logIndex++
		debugP := fmt.Sprintf("P:%02X", debugger.CpuStatus)
		if p != debugP {
			fmt.Printf("Status: %08b\n", debugger.CpuStatus)
			t.Errorf("Expected p %s, got %s\n", p, debugP)
		}

		sp := logData[logIndex]
		logIndex++
		debugSP := fmt.Sprintf("SP:%02X", debugger.CpuSP)
		if sp != debugSP {
			t.Fatalf("Expected sp %s, got %s\n", sp, debugSP)
		}

		cycles := logData[logIndex]

		debugCycles := fmt.Sprintf("CYC:%d", debugger.CpuCycles)
		if debugCycles != cycles {
			t.Fatalf("Expected cycles %s, got %s\n", cycles, debugCycles)
		}

		debugger.Reset()
	}
}
