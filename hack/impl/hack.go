package impl

import (
	"fmt"
)

// Hack - Struct representing the hack computer architecture
type Hack struct {
	rom       []uint16 // Instruction Memory: 32k ROM (0-32767)
	ram       []int16  // Data Memory: 16k RAM (0-16383), 8k screen map (16384-24575), 16 bit keyboard register (24576)
	aRegister int16    // 16 bit A register
	dRegister int16    // 16 bit D register
	pc        int16    // 16 bit Program Counter
}

// NewHack - Constructor
func NewHack() *Hack {
	h := Hack{}
	h.rom = make([]uint16, 32768, 32768)
	h.ram = make([]int16, 24577, 24577)
	return &h
}

// LoadRom - Loads a program into the instruction memory
func (h *Hack) LoadRom(data []uint16) {
	copy(h.rom[:], data)
}

// SetKeyboard - Loads a key code into the keyboard register
func (h *Hack) SetKeyboard(data int16) {
	h.ram[24576] = data
}

// GetScreen - Returns the screen memory map data
func (h *Hack) GetScreen() []int16 {
	return h.ram[16384:24576]
}

// Tick - Simulates one CPU cycle
func (h *Hack) Tick() {
	inst := int16(h.rom[h.pc])
	fmt.Printf("A=%5d | D=%5d | PC=%2d | KEY=%3d | %016b | ", h.aRegister, h.dRegister, h.pc, h.ram[24576], h.rom[h.pc])
	h.pc++
	if (inst>>15)&0b1 == 0 {
		// Execute A-Instruction (0vvvvvvvvvvvvvvv: aRegister=vvvvvvvvvvvvvvv)
		h.aRegister = inst
	} else {
		// Execute C-Instruction (111accccccdddjjj: comp=acccccc, dest=ddd, jump=jjj)
		comp := (inst >> 6) & 0b1111111
		dest := (inst >> 3) & 0b111
		jump := inst & 0b111
		fmt.Printf("comp=%07b ", comp)
		fmt.Printf("dest=%03b ", dest)
		fmt.Printf("jump=%03b ", jump)
		computed, _ := h.compute(comp)

		if (dest>>2)&0b1 == 1 {
			h.aRegister = computed
		}
		if (dest>>1)&0b1 == 1 {
			h.dRegister = computed
		}
		if dest&0b1 == 1 {
			h.ram[h.aRegister] = computed
		}

		h.handleJump(jump, computed)
	}
	fmt.Print("\n")
}

// Reset - Resets the computer
func (h *Hack) Reset() {
	for i := range h.ram {
		h.ram[i] = 0b0000000000000000
	}
	h.aRegister = 0b0000000000000000
	h.dRegister = 0b0000000000000000
	h.pc = 0b0000000000000000
}

func (h *Hack) compute(comp int16) (int16, error) {
	switch comp {
	case 0b0101010:
		fmt.Print("comp(0)   ")
		return 0, nil
	case 0b0111111:
		fmt.Print("comp(1)   ")
		return 1, nil
	case 0b0111010:
		fmt.Print("comp(-1)  ")
		return -1, nil
	case 0b0001100:
		fmt.Print("comp(D)   ")
		return h.dRegister, nil
	case 0b0110000:
		fmt.Print("comp(A)   ")
		return h.aRegister, nil
	case 0b0001101:
		fmt.Print("comp(!D)  ")
		return ^h.dRegister, nil // bitwise not
	case 0b0110001:
		fmt.Print("comp(!A)  ")
		return ^h.aRegister, nil // bitwise not
	case 0b0001111:
		fmt.Print("comp(-D)  ")
		return -h.dRegister, nil
	case 0b0110011:
		fmt.Print("comp(-A)  ")
		return -h.aRegister, nil
	case 0b0011111:
		fmt.Print("comp(D+1) ")
		return h.dRegister + 1, nil
	case 0b0110111:
		fmt.Print("comp(A+1) ")
		return h.aRegister + 1, nil
	case 0b0001110:
		fmt.Print("comp(D-1) ")
		return h.dRegister - 1, nil
	case 0b0110010:
		fmt.Print("comp(A-1) ")
		return h.aRegister - 1, nil
	case 0b0000010:
		fmt.Print("comp(D+A) ")
		return h.dRegister + h.aRegister, nil
	case 0b0010011:
		fmt.Print("comp(D-A) ")
		return h.dRegister - h.aRegister, nil
	case 0b0000111:
		fmt.Print("comp(A-D) ")
		return h.aRegister - h.dRegister, nil
	case 0b0000000:
		fmt.Print("comp(D&A) ")
		return h.dRegister & h.aRegister, nil
	case 0b0010101:
		fmt.Print("comp(D|A) ")
		return h.dRegister | h.aRegister, nil
	case 0b1110000:
		fmt.Print("comp(M)   ")
		return h.ram[h.aRegister], nil
	case 0b1110001:
		fmt.Print("comp(!M)  ")
		return ^h.ram[h.aRegister], nil // bitwise not
	case 0b1110011:
		fmt.Print("comp(-M)  ")
		return -h.ram[h.aRegister], nil
	case 0b1110111:
		fmt.Print("comp(M+1) ")
		return h.ram[h.aRegister] + 1, nil
	case 0b1110010:
		fmt.Print("comp(M-1) ")
		return h.ram[h.aRegister] - 1, nil
	case 0b1000010:
		fmt.Print("comp(D+M) ")
		return h.dRegister + h.ram[h.aRegister], nil
	case 0b1010011:
		fmt.Print("comp(D-M) ")
		return h.dRegister - h.ram[h.aRegister], nil
	case 0b1000111:
		fmt.Print("comp(M-D) ")
		return h.ram[h.aRegister] - h.dRegister, nil
	case 0b1000000:
		fmt.Print("comp(D&M) ")
		return h.dRegister & h.ram[h.aRegister], nil
	case 0b1010101:
		fmt.Print("comp(D|M) ")
		return h.dRegister | h.ram[h.aRegister], nil
	default:
		fmt.Print("comp(ERR) ")
		return 0, fmt.Errorf("Invalid comp: %07b", comp)
	}
}

func (h *Hack) handleJump(jump, computed int16) {
	switch jump {
	case 0b000:
		fmt.Print("jump(---)")
	case 0b001:
		fmt.Print("jump(JGT)")
		if computed > 0 {
			h.pc = h.aRegister
		}
	case 0b010:
		fmt.Print("jump(JEQ)")
		if computed == 0 {
			h.pc = h.aRegister
		}
	case 0b011:
		fmt.Print("jump(JGE)")
		if computed >= 0 {
			h.pc = h.aRegister
		}
	case 0b100:
		fmt.Print("jump(JLT)")
		if computed < 0 {
			h.pc = h.aRegister
		}
	case 0b101:
		fmt.Print("jump(JNE)")
		if computed != 0 {
			h.pc = h.aRegister
		}
	case 0b110:
		fmt.Print("jump(JLE)")
		if computed <= 0 {
			h.pc = h.aRegister
		}
	case 0b111:
		fmt.Print("jump(JMP)")
		h.pc = h.aRegister
	}
}
