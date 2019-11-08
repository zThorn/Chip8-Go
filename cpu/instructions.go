package cpu

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

func (cpu *C8CPU) op_6XNN(opcode uint16) {
	cpu.registers[cpu.vx]  = opcode & 0x00ff
}

func (cpu *C8CPU) op_7XNN(opcode uint16) {
	cpu.registers[cpu.vx]  += opcode & 0x00ff
}

func (cpu *C8CPU) op_8XY0(opcode uint16) {
	cpu.registers[cpu.vx]  = cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY1(opcode uint16) {
	cpu.registers[cpu.vx]  = cpu.registers[cpu.vx] | cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY2(opcode uint16) {
	cpu.registers[cpu.vx]  = cpu.registers[cpu.vx] & cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY3(opcode uint16) {
	cpu.registers[cpu.vx]  = cpu.registers[cpu.vx] ^ cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY4(opcode uint16) {
	if cpu.GetVXR()+cpu.GetVYR() > 0xff{
		cpu.SetVF(1)
	} else{
		cpu.SetVF(0)
	}
	cpu.registers[cpu.vx]  += cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY5(opcode uint16) {
	if cpu.GetVXR() < cpu.GetVYR(){
		cpu.SetVF(0)
	} else{
		cpu.SetVF(1)
	}
	cpu.registers[cpu.vx]  -= cpu.registers[cpu.vy]
}

func (cpu *C8CPU) op_8XY6(opcode uint16) {
	cpu.SetVF(cpu.registers[cpu.vx] & 0x01)
	cpu.SetRegVX(cpu.GetVXR() >> 1)
}

func (cpu *C8CPU) op_8XY7(opcode uint16) {
	if cpu.GetVYR() < cpu.GetVXR(){
		cpu.SetVF(0)
	} else{
		cpu.SetVF(1)
	}
	cpu.registers[cpu.vy] -= cpu.GetVYR()
}
