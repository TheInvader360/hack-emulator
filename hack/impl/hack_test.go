package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	type test struct {
		comp, expectedComputed int16
		expectedErrMsg         string
	}
	tests := []test{
		{comp: 0b0101010, expectedComputed: 0, expectedErrMsg: ""},      // 0
		{comp: 0b0111111, expectedComputed: 1, expectedErrMsg: ""},      // 1
		{comp: 0b0111010, expectedComputed: -1, expectedErrMsg: ""},     // -1
		{comp: 0b0001100, expectedComputed: 4369, expectedErrMsg: ""},   // D
		{comp: 0b0110000, expectedComputed: 13107, expectedErrMsg: ""},  // A
		{comp: 0b0001101, expectedComputed: -4370, expectedErrMsg: ""},  // !D i.e. bitwise not of 0b0001000100010001 = 0b1110111011101110 = decimal -4370 (two's complement)
		{comp: 0b0110001, expectedComputed: -13108, expectedErrMsg: ""}, // !A i.e. bitwise not of 0b0011001100110011 = 0b1100110011001100 = decimal -13108 (two's complement)
		{comp: 0b0001111, expectedComputed: -4369, expectedErrMsg: ""},  // -D
		{comp: 0b0110011, expectedComputed: -13107, expectedErrMsg: ""}, // -A
		{comp: 0b0011111, expectedComputed: 4370, expectedErrMsg: ""},   // D+1
		{comp: 0b0110111, expectedComputed: 13108, expectedErrMsg: ""},  // A+1
		{comp: 0b0001110, expectedComputed: 4368, expectedErrMsg: ""},   // D-1
		{comp: 0b0110010, expectedComputed: 13106, expectedErrMsg: ""},  // A-1
		{comp: 0b0000010, expectedComputed: 17476, expectedErrMsg: ""},  // D+A
		{comp: 0b0010011, expectedComputed: -8738, expectedErrMsg: ""},  // D-A
		{comp: 0b0000111, expectedComputed: 8738, expectedErrMsg: ""},   // A-D
		{comp: 0b0000000, expectedComputed: 4369, expectedErrMsg: ""},   // D&A i.e. 0b0001000100010001 & 0b0011001100110011 = 0b0001000100010001 = decimal 4369
		{comp: 0b0010101, expectedComputed: 13107, expectedErrMsg: ""},  // D|A i.e. 0b0001000100010001 | 0b0011001100110011 = 0b0011001100110011 = decimal 13107
		{comp: 0b1110000, expectedComputed: 1799, expectedErrMsg: ""},   // M
		{comp: 0b1110001, expectedComputed: -1800, expectedErrMsg: ""},  // !M i.e. bitwise not of 0b0000011100000111 = 0b1111100011111000 = decimal -1800 (two's complement)
		{comp: 0b1110011, expectedComputed: -1799, expectedErrMsg: ""},  // -M
		{comp: 0b1110111, expectedComputed: 1800, expectedErrMsg: ""},   // M+1
		{comp: 0b1110010, expectedComputed: 1798, expectedErrMsg: ""},   // M-1
		{comp: 0b1000010, expectedComputed: 6168, expectedErrMsg: ""},   // D+M
		{comp: 0b1010011, expectedComputed: 2570, expectedErrMsg: ""},   // D-M
		{comp: 0b1000111, expectedComputed: -2570, expectedErrMsg: ""},  // M-D
		{comp: 0b1000000, expectedComputed: 257, expectedErrMsg: ""},    // D&M i.e. 0b0001000100010001 & 0b0000011100000111 = 0b0000000100000001 = decimal 257
		{comp: 0b1010101, expectedComputed: 5911, expectedErrMsg: ""},   // D|M i.e. 0b0001000100010001 | 0b0000011100000111 = 0b0001011100010111 = decimal 5911
		{comp: 0b1101010, expectedComputed: 0, expectedErrMsg: "Invalid comp: 1101010"},
	}
	for _, tc := range tests {
		h := NewHack()
		h.aRegister = 13107       // i.e. 0b0011001100110011
		h.dRegister = 4369        // i.e. 0b0001000100010001
		h.ram[h.aRegister] = 1799 // i.e. 0b0000011100000111
		computed, err := h.compute(tc.comp)
		assert.Equal(t, tc.expectedComputed, computed)
		if len(tc.expectedErrMsg) > 0 {
			assert.EqualError(t, err, tc.expectedErrMsg)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestHandleJump(t *testing.T) {
	type test struct {
		jump, computed, expectedPC int16
	}
	tests := []test{
		{jump: 0b000, computed: 0, expectedPC: 0b100},  // no jump
		{jump: 0b001, computed: 1, expectedPC: 0b111},  // jump if computed > 0 (jump)
		{jump: 0b001, computed: 0, expectedPC: 0b100},  // jump if computed > 0 (no jump)
		{jump: 0b010, computed: 0, expectedPC: 0b111},  // jump if computed == 0 (jump)
		{jump: 0b010, computed: 1, expectedPC: 0b100},  // jump if computed == 0 (no jump)
		{jump: 0b011, computed: 1, expectedPC: 0b111},  // jump if computed >= 0 (jump)
		{jump: 0b011, computed: 0, expectedPC: 0b111},  // jump if computed >= 0 (jump)
		{jump: 0b011, computed: -1, expectedPC: 0b100}, // jump if computed >= 0 (no jump)
		{jump: 0b100, computed: -1, expectedPC: 0b111}, // jump if computed < 0 (jump)
		{jump: 0b100, computed: 0, expectedPC: 0b100},  // jump if computed < 0 (no jump)
		{jump: 0b101, computed: 1, expectedPC: 0b111},  // jump if computed != 0 (jump)
		{jump: 0b101, computed: 0, expectedPC: 0b100},  // jump if computed != 0 (no jump)
		{jump: 0b110, computed: -1, expectedPC: 0b111}, // jump if computed <= 0 (jump)
		{jump: 0b110, computed: 0, expectedPC: 0b111},  // jump if computed <= 0 (jump)
		{jump: 0b110, computed: 1, expectedPC: 0b100},  // jump if computed <= 0 (no jump)
		{jump: 0b111, computed: 0, expectedPC: 0b111},  // jump
	}
	for _, tc := range tests {
		h := NewHack()
		h.aRegister = 0b111
		h.pc = 0b100
		h.handleJump(tc.jump, tc.computed)
		assert.Equal(t, tc.expectedPC, h.pc)
	}
}
