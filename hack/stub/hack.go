package stub

import "fmt"

// Hack - Stub Struct
type Hack struct {
	screenStateFlag bool
}

// NewHack - Constructor
func NewHack() *Hack {
	h := Hack{}
	return &h
}

// LoadRom - Stub
func (h *Hack) LoadRom(data []uint16) {
	fmt.Println("LoadRom()", data)
}

// SetKeyboard - Stub
func (h *Hack) SetKeyboard(data int16) {
	if data > 0 {
		h.screenStateFlag = false
	} else {
		h.screenStateFlag = true
	}
}

// GetScreen - Stub
func (h *Hack) GetScreen() []int16 {
	pixels := []int16{}
	for i := 0; i < 256*32; i++ {
		if i%2 == 0 && h.screenStateFlag && i < 256*16 {
			pixels = append(pixels, -1)
		} else {
			pixels = append(pixels, 0)
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
