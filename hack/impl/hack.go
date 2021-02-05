package impl

import "fmt"

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

	// Temp data (stripey screen)
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
	fmt.Println(h.rom)
}

// SetKeyboard - Loads a key code into the keyboard register
func (h *Hack) SetKeyboard(data uint16) {
	h.ram[24576] = data
	fmt.Println(fmt.Sprintf("KEY=%d | PC=%d", h.ram[24576], h.pc))
}

// GetScreen - Returns the screen memory map data
func (h *Hack) GetScreen() []uint16 {
	return h.ram[16384:24576]
}

// Tick - Simulates one CPU cycle
func (h *Hack) Tick() {
	//TODO
	h.pc++
}

// Reset - Resets the program counter
func (h *Hack) Reset() {
	h.pc = 0
}
