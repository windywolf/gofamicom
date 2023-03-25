package main

import (
	"fmt"
	"gofamicom/apu"
	"gofamicom/cartridge"
	"gofamicom/cpu"
	"gofamicom/ppu"
	"os"
)

type Famicom struct {
	cpu     *cpu.CPU
	ppu     *ppu.PPU
	apu     *apu.APU
	cart    *cartridge.Cartridge
	ram     []byte
	prgBank []byte
	chrBank []byte
	cycles  uint64
}

func (f *Famicom) reset() {
	// TODO: reset the Famicom
}

func (f *Famicom) run() {
	// TODO: run the Famicom
}

func (f *Famicom) read(addr uint16) byte {
	// TODO: read from the CPU memory map

	return 0
}

func (f *Famicom) write(addr uint16, data byte) {
	// TODO: write to the CPU memory map
}

func (f *Famicom) step() {
	// TODO: execute one CPU instruction
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gofamicom <path-to-nes-rom>")
		return
	}

	f := &Famicom{}
	f.cpu = cpu.NewCPU()
	f.ppu = ppu.NewPPU()
	f.apu = apu.NewAPU()
	f.cart = cartridge.NewCartridge(os.Args[1])
	f.ram = make([]byte, 2048)
	f.prgBank = make([]byte, 0x4000)
	f.chrBank = make([]byte, 0x2000)

	f.reset()

	for {
		f.step()
		f.ppu.Render()
	}
}
