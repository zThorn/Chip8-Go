package cpu

import "fmt"

type Instruction struct {
	Opcode  uint16
	Execute func(uint16)
}

//Return
func (cpu *C8CPU) op_1NNN(opcode uint16) {
	cpu.PC = opcode & 0x0fff
}

func (cpu *C8CPU) op_2NNN(opcode uint16) {
	cpu.stack = append(cpu.stack, cpu.PC)
	cpu.PC = opcode & 0x0fff
}

func (cpu *C8CPU) op_3XNN(opcode uint16) {
	if cpu.vxr == (opcode & 0x00ff) {
		cpu.PC += 2
	}
}

func (cpu *C8CPU) op_4XNN(opcode uint16) {
	if cpu.vxr != (opcode & 0x00ff) {
		cpu.PC += 2
	}
}

func (cpu *C8CPU) op_5XY0(opcode uint16) {
	if cpu.vxr == cpu.vyr {
		cpu.PC += 2
	}
}
