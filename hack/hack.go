package hack

//Hack - Interface
type Hack interface {
	LoadRom([]uint16)
	SetKeyboard(int16)
	GetScreen() []int16
	Tick()
	Reset()
}
