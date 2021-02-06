package impl

import (
	"fmt"
	"os"
)

// Hack - Struct representing the hack computer architecture
type Hack struct {
	rom       []uint16 // Instruction Memory: 32k ROM (0-32767)
	ram       []uint16 // Data Memory: 16k RAM (0-16383), 8k screen map (16384-24575), 16 bit keyboard register (24576)
	aRegister uint16
	dRegister uint16
	pc        uint16
}

// NewHack - Constructor
func NewHack() *Hack {
	h := Hack{}
	h.rom = make([]uint16, 32768, 32768)
	h.ram = make([]uint16, 24577, 24577)

	// TODO - Remove (temp)
	for i, _ := range h.ram {
		if i%2 == 0 || i%3 == 0 {
			h.ram[i] = 0b0000000000000000
		} else {
			h.ram[i] = 0b1111111111111111
		}
	}

	return &h
}

// LoadRom - Loads a program into the instruction memory
func (h *Hack) LoadRom(data []uint16) {
	copy(h.rom[:], data)
}

// SetKeyboard - Loads a key code into the keyboard register
func (h *Hack) SetKeyboard(data uint16) {
	h.ram[24576] = data
}

// GetScreen - Returns the screen memory map data
func (h *Hack) GetScreen() []uint16 {
	return h.ram[16384:24576]
}

// Tick - Simulates one CPU cycle
func (h *Hack) Tick() {
	h.handleInstruction()
	h.pc++

	// TODO - Remove (temp)
	if h.pc > 40 {
		os.Exit(0)
	}
}

// Reset - Resets the program counter
func (h *Hack) Reset() {
	h.pc = 0
}

func (h *Hack) handleInstruction() {
	inst := h.rom[h.pc]
	fmt.Printf("A=%5d | D=%5d | PC=%2d | KEY=%3d | %016b | ", h.aRegister, h.dRegister, h.pc, h.ram[24576], inst)
	if (inst>>15)&0b1 == 0 {
		h.aRegister = inst
		fmt.Print("A |\n")
	} else {
		//  111accccccdddjjj
		// c   acccccc
		// d          ddd
		// j             jjj
		c := (inst >> 6) & 0b1111111
		d := (inst >> 3) & 0b111
		j := inst & 0b111
		fmt.Print("C | ")
		fmt.Printf("c=%07b ", c)
		fmt.Printf("d=%03b ", d)
		fmt.Printf("j=%03b\n", j)
		h.executeComputeInstruction(c, d, j)
	}
}

func (h *Hack) executeComputeInstruction(c, d, j uint16) {
	// TODO - Implement...
}
