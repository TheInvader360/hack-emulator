package stub

import "fmt"

// Hack - Stub Struct
type Hack struct {
	screenStateFlag bool
}

// LoadRom - Stub
func (h *Hack) LoadRom(data []uint16) {
	fmt.Println("LoadRom()", data)
}

// SetKeyboard - Stub
func (h *Hack) SetKeyboard(data uint16) {
	if data > 0 {
		h.screenStateFlag = false
	} else {
		h.screenStateFlag = true
	}
}

// GetScreen - Stub
func (h *Hack) GetScreen() []uint16 {
	pixels := []uint16{}
	for y := 0; y < 256; y++ {
		for x := 0; x < 32; x++ {
			if x%2 == 0 && h.screenStateFlag {
				pixels = append(pixels, 0b1111111111111111)
			} else {
				pixels = append(pixels, 0b0000000000000000)
			}
		}
	}
	return pixels
}

// Tick - Stub
func (h *Hack) Tick() {
	fmt.Println("Tick()")
}

// Reset - Stub
func (h *Hack) Reset() {
	fmt.Println("Reset()")
}
