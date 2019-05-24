package cpu

type C8CPU struct {
	fonts        Fonts
	registers    [16]uint16
	memory       [4096]uint16
	PC           uint16
	stack        []uint16
	vx           uint16
	vy           uint16
	vxr          uint16
	vyr          uint16
	Instructions map[uint16]func(uint16)
	debug        bool
}

type font struct {
	hexcodes    []byte
	description string
}

type Fonts struct {
	fonts []font
}

func NewCPU(debug bool) *C8CPU {
	cpu := new(C8CPU)
	cpu.stack = make([]uint16, 16, 16)
	cpu.Instructions = make(map[uint16]func(uint16))
	cpu.loadFonts()
	cpu.loadInstructions()
	cpu.debug = debug
	return cpu

}

func (cpu *C8CPU) loadInstructions() {
	cpu.Instructions[0x1000] = cpu.op_1NNN
	cpu.Instructions[0x2000] = cpu.op_2NNN
	cpu.Instructions[0x3000] = cpu.op_3XNN
	cpu.Instructions[0x4000] = cpu.op_4XNN
	cpu.Instructions[0x5000] = cpu.op_5XY0
}

func (cpu *C8CPU) Cycle(opcode uint16) {
	//var opcode uint16 = (cpu.memory[cpu.pc] << uint(8)) | (cpu.memory[cpu.pc+1])
	instr := opcode & 0xf000
	//Advance the PC
	cpu.PC += 2
	if cpu.debug != true {

		cpu.vx = (opcode & 0x0f00)
		cpu.vx = cpu.vx >> 8
		cpu.vy = (opcode & 0x00f0) >> uint(4)
	}
	cpu.vxr = cpu.registers[cpu.vx]
	cpu.vyr = cpu.registers[cpu.vy]
	cpu.Instructions[instr](opcode)
}
func (cpu *C8CPU) SetVX(vx uint16) {
	cpu.vx = vx
}
func (cpu *C8CPU) SetVY(vy uint16) {
	cpu.vy = vy
}

func (cpu *C8CPU) SetRegVX(val uint16) {
	cpu.registers[cpu.vx] = val
}

func (cpu *C8CPU) SetRegVXVY(vx uint16, vy uint16, val uint16) {
	cpu.registers[cpu.vx] = val
	cpu.registers[cpu.vy] = val
}

//Will literally just clear out all registers
func (cpu *C8CPU) ClearCPU() {

	print("We compiled!")
}

func (cpu *C8CPU) popStack() {
	n := len(cpu.stack) - 1
	cpu.stack[n] = 0

	cpu.stack = cpu.stack[:n]
}

func (cpu *C8CPU) loadFonts() {
	cpu.fonts = Fonts{[]font{
		{[]byte{0xF0, 0x90, 0x90, 0x90, 0xF0}, "0"},
		{[]byte{0x20, 0x60, 0x20, 0x20, 0x70}, "1"},
		{[]byte{0xF0, 0x10, 0xF0, 0x80, 0xF0}, "2"},
		{[]byte{0xF0, 0x10, 0xF0, 0x10, 0xF0}, "3"},
		{[]byte{0x90, 0x90, 0xF0, 0x10, 0x10}, "4"},
		{[]byte{0xF0, 0x80, 0xF0, 0x10, 0xF0}, "5"},
		{[]byte{0xF0, 0x80, 0xF0, 0x90, 0xF0}, "6"},
		{[]byte{0xF0, 0x10, 0x20, 0x40, 0x40}, "7"},
		{[]byte{0xF0, 0x90, 0xF0, 0x90, 0xF0}, "8"},
		{[]byte{0xF0, 0x90, 0xF0, 0x10, 0xF0}, "9"},
		{[]byte{0xF0, 0x90, 0xF0, 0x90, 0x90}, "A"},
		{[]byte{0xE0, 0x90, 0xE0, 0x90, 0xE0}, "B"},
		{[]byte{0xF0, 0x80, 0x80, 0x80, 0xF0}, "C"},
		{[]byte{0xE0, 0x90, 0x90, 0x90, 0xE0}, "D"},
		{[]byte{0xF0, 0x80, 0xF0, 0x80, 0xF0}, "E"},
		{[]byte{0xF0, 0x80, 0xF0, 0x80, 0x80}, "F"},
	}}
}
