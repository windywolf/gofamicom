package apu

type APU struct {
	// TODO: APU fields
}

func NewAPU() *APU {
	a := &APU{}
	// TODO: APU initialization
	return a
}

func (a *APU) Reset() {
	// TODO: APU reset
}

func (a *APU) Step() {
	// TODO: APU step
}

func (a *APU) Read(addr uint16) byte {
	// TODO: APU read
	return 0
}

func (a *APU) Write(addr uint16, data byte) {
	// TODO: APU write
}
