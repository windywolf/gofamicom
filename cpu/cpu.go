package cpu

type CPU struct {
	// TODO: CPU fields
}

func NewCPU() *CPU {
	c := &CPU{}
	// TODO: CPU initialization
	return c
}

func (c *CPU) Reset() {
	// TODO: CPU reset
}

func (c *CPU) Step() {
	// TODO: CPU step
}

func (c *CPU) Read(addr uint16) byte {
	// TODO: CPU read

	return 0
}

func (c *CPU) Write(addr uint16, data byte) {
	// TODO: CPU write
}
