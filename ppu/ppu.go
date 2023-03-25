package ppu

type PPU struct {
	// TODO: PPU fields
}

func NewPPU() *PPU {
	p := &PPU{}
	// TODO: PPU initialization
	return p
}

func (p *PPU) Reset() {
	// TODO: PPU reset
}

func (p *PPU) Step() {
	// TODO: PPU step
}

func (p *PPU) Render() {
	// TODO: PPU rendering
}

func (p *PPU) Read(addr uint16) byte {
	// TODO: PPU read

	return 0
}

func (p *PPU) Write(addr uint16, data byte) {
	// TODO: PPU write
}
