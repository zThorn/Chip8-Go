package cpu

type Instruction struct {
	Opcode      byte
	Description string
	Execute     func(cpu *C8CPU)
}
