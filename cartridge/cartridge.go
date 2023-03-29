package cartridge

import (
	"io/ioutil"
)

const (
	PRGROMBankSize = 16 * 1024
	CHRROMBankSize = 8 * 1024
)

type Cartridge struct {
	prgROM []byte
	chrROM []byte
}

func loadNESROM(romData []byte) (*Cartridge, error) {
	// Parse header
	prgROMSize := int(romData[4]) * PRGROMBankSize
	chrROMSize := int(romData[5]) * CHRROMBankSize

	// Load PRG ROM
	prgROMStart := 16
	prgROMEnd := prgROMStart + prgROMSize
	prgROM := romData[prgROMStart:prgROMEnd]

	// Load CHR ROM
	chrROMStart := prgROMEnd
	chrROMEnd := chrROMStart + chrROMSize
	chrROM := romData[chrROMStart:chrROMEnd]

	// Return cartridge
	return &Cartridge{
		prgROM: prgROM,
		chrROM: chrROM,
	}, nil
}

func NewCartridge(filename string) *Cartridge {
	romData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	c, _ := loadNESROM(romData)
	// TODO: Cartridge initialization
	return c
}

func (c *Cartridge) Load(filename string) {
	// TODO: Cartridge loading
}

func (c *Cartridge) Read(addr uint16) byte {
	// TODO: Cartridge read

	return 0
}

func (c *Cartridge) Write(addr uint16, data byte) {
	// TODO: Cartridge write
}
