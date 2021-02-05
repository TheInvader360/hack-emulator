package hack

//Hack - Interface
type Hack interface {
	LoadRom([]uint16)
	SetKeyboard(uint16)
	GetScreen() []uint16
	Tick()
	Reset()
}
