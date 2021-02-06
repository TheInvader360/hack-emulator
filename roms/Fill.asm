  // first = SCREEN (or screen's base address of 16384)
  @SCREEN
  D=A
  @first
  M=D
  // finished = SCREEN + 32 * 256 (or screen's last address + 1)
  @8192
  D=D+A
  @finished
  M=D

(MAIN_LOOP)
  // i = first
  @first
  D=M
  @i
  M=D
  // if input = 0 goto OFF
  @KBD
  D=M
  @OFF
  D;JEQ
  // else goto ON
  @ON
  0;JMP

(OFF)
  // state = 0 (i.e. 0000000000000000)
  @state
  M=0
  // goto REFRESH
  @REFRESH
  0;JMP

(ON)
  // state = -1 (i.e. 1111111111111111 in two's complement)
  @state
  M=-1
  // goto REFRESH
  @REFRESH
  0;JMP

(REFRESH)
  // RAM[i] = state
  @state
  D=M                    // D = state (0 or -1 / 0000000000000000 or 1111111111111111)
  @i
  A=M                    // A = i
  M=D                    // RAM[i] = state
  // i = i + 1
  @i
  M=M+1
  // if i = finished goto MAIN_LOOP
  @i
  D=M                    // D = i
  @finished
  D=D-M                  // D = i - finished
  @MAIN_LOOP
  D;JEQ                  // if i - finished = 0 goto MAIN_LOOP (equivalent to: if i = finished goto MAIN_LOOP)
  // else goto REFRESH
  @REFRESH
  0;JMP                  // goto REFRESH
