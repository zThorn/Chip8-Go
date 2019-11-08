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
	print("Cycling")
	c.Cycle(0x5022)
	if c.PC != 4 {
		t.Errorf("PC was incorrect, got %d expected %d", c.PC, 4)
		return
	}

	fmt.Printf("5XY0 Succeeded!\n")
}

func Test6XNN(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)


	c.Cycle(0x6033)

	if c.GetVXR() != 0x33 {
		t.Errorf("VX was incorrect, got %d expected %d", c.GetVX(), 0x33)
		return
	}

	fmt.Printf("6XNN Succeeded!\n")
}

func Test7XNN(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)


	c.Cycle(0x6033)

	if c.GetVXR() != 0x33 {
		t.Errorf("VX was incorrect, got %d expected %d", c.GetVX(), 0x33)
		return
	}

	fmt.Printf("7XNN Succeeded!\n")
}

func Test8XY0(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)
	c.SetVY(1)

	c.SetRegVY(0x25)
	c.Cycle(0x8030)

	if c.GetVXR() != 0x25 {
		t.Errorf("VXR was incorrect, got %d expected %d", c.GetVXR(), 0x25)
		return
	}

	fmt.Printf("8XY0 Succeeded!\n")
}

func Test8XY1(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)
	c.SetVY(1)

	c.SetRegVX(0x20)
	c.SetRegVY(0x25)
	c.Cycle(0x8031)

	if c.GetVXR() != 0x25 {
		t.Errorf("VXR was incorrect, got %d expected %d", c.GetVXR(), 0x25)
		return
	}

	fmt.Printf("8XY1 Succeeded!\n")
}

func Test8XY2(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)
	c.SetVY(1)

	c.SetRegVX(0x20)
	c.SetRegVY(0x25)
	c.Cycle(0x8032)

	if c.GetVXR() != 0x20 {
		t.Errorf("VXR was incorrect, got %d expected %d", c.GetVXR(), 0x20)
		return
	}

	fmt.Printf("8XY2 Succeeded!\n")
}

func Test8XY3(t *testing.T) {

	c := cpu.NewCPU(true)
	c.SetVX(0)
	c.SetVY(1)

	c.SetRegVX(0x20)
	c.SetRegVY(0x25)
	c.Cycle(0x8030)

	if c.GetVXR() == 0x20 {
		t.Errorf("VXR was incorrect, got %d expected %d", c.GetVXR(), 0x25)
		return
	}

	fmt.Printf("8XY2 Succeeded!\n")
}