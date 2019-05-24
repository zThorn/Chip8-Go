package main

import "github.com/zThorn/Gate/cpu"
import "testing"
import "fmt"

func Test1NNN(t *testing.T) {
	c := cpu.NewCPU(true)
	c.Cycle(0x1222)
	if c.PC != 0x222 {
		t.Errorf("PC was incorrect, got %d expected %d", c.PC, 0x222)
		return
	}

	fmt.Printf("1NNN Succeeded!\n")
}

func Test3NNN(t *testing.T) {
	c := cpu.NewCPU(true)
	c.SetRegVX(0x22)
	c.Cycle(0x3022)

	if c.PC != 4 {
		t.Errorf("PC was incorrect, got %d expected %d", c.PC, 4)
		return
	}

	fmt.Printf("3NNN Succeeded!\n")
}

func Test4NNN(t *testing.T) {
	c := cpu.NewCPU(true)
	c.SetRegVX(0x21)
	c.Cycle(0x3022)

	if c.PC != 2 {
		t.Errorf("PC was incorrect, got %d expected %d", c.PC, 2)
		return
	}

	fmt.Printf("4NNN Succeeded!\n")
}

func Test5XY0(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)
	c.SetVY(1)
	c.SetRegVXVY(0, 1, 0x21)

	c.Cycle(0x5022)
	if c.PC != 4 {
		t.Errorf("PC was incorrect, got %d expected %d", c.PC, 4)
		return
	}

	fmt.Printf("4NNN Succeeded!\n")
}
