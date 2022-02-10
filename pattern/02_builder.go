package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Computer struct
type Computer struct {
	systemBlock     string
	motherBoard     string
	cpu             string
	gpu             string
	powerSypplyUnit string
	ram             string
	fan             string
}

// ComputerBuilder interface for all builder types
type ComputerBuilder interface {
	setSystemBlock()
	setMotherBoard()
	setCPU()
	setGPU()
	setPowerSupplyUnit()
	setRAM()
	setFAN()
	getComputer() *Computer
}

// getBuilder - return specific Builder Type
func getBuilder(builderType string) ComputerBuilder {
	if builderType == "office" {
		return newOfficeComputerBuilder()
	}

	if builderType == "gaming" {
		return newGamingComputerBuilder()
	}

	return nil
}

// officeComputerBuilder - contains basic computer elements for usuall work
type officeComputerBuilder struct {
	systemBlock     string
	motherBoard     string
	cpu             string
	gpu             string
	powerSypplyUnit string
	ram             string
	fan             string
}

// newOfficeComputerBuilder - creates new BasicComputerBilder instance
func newOfficeComputerBuilder() *officeComputerBuilder {
	return &officeComputerBuilder{}
}

// setSystemBlock - sets simple system block for basic computer
func (b *officeComputerBuilder) setSystemBlock() {
	b.systemBlock = "Some standard white system block"
}

// setMotherBoard - sets simple motherboard for basic computer
func (b *officeComputerBuilder) setMotherBoard() {
	b.motherBoard = "Some simple motherboard"
}

// setCPU - sets simple cpu for basic computer
func (b *officeComputerBuilder) setCPU() {
	b.cpu = "Some simple cpu"
}

// setGPU - sets simple gpu for basic computer
func (b *officeComputerBuilder) setGPU() {
	b.gpu = "Some simple gpu"
}

// setPowerSupplyUnit - sets simple powerSupplyUnit for basic computer
func (b *officeComputerBuilder) setPowerSupplyUnit() {
	b.powerSypplyUnit = "Some simple powerSupplyUnit"
}

// setRAM - sets simple RAM for basic computer
func (b *officeComputerBuilder) setRAM() {
	b.ram = "Some simple RAM"
}

// setFAN - sets simple FAN for basic computer
func (b *officeComputerBuilder) setFAN() {
	b.fan = "Some simple FAN"
}

// getComputer - returns simple computer
func (b *officeComputerBuilder) getComputer() *Computer {
	return &Computer{
		motherBoard:     b.motherBoard,
		systemBlock:     b.systemBlock,
		cpu:             b.cpu,
		gpu:             b.gpu,
		powerSypplyUnit: b.powerSypplyUnit,
		ram:             b.ram,
		fan:             b.fan,
	}
}

// gamingComputerBuilder - contains gaming computer elements for best gaming experince
type gamingComputerBuilder struct {
	systemBlock     string
	motherBoard     string
	cpu             string
	gpu             string
	powerSypplyUnit string
	ram             string
	fan             string
}

// newGamingComputerBuilder - creates new GamingComputerBuilder instance
func newGamingComputerBuilder() *gamingComputerBuilder {
	return &gamingComputerBuilder{}
}

// setSystemBlock - sets gaming system block for gaming computer
func (b *gamingComputerBuilder) setSystemBlock() {
	b.systemBlock = "Some gaming black with window system block"
}

// setMotherBoard - sets gaming motherboard for gaming computer
func (b *gamingComputerBuilder) setMotherBoard() {
	b.motherBoard = "Some gaming motherboard"
}

// setCPU - sets gaming cpu for gaming computer
func (b *gamingComputerBuilder) setCPU() {
	b.cpu = "Some gaming cpu"
}

// setGPU - sets gaming gpu for gaming computer
func (b *gamingComputerBuilder) setGPU() {
	b.gpu = "Some gaming gpu"
}

// setPowerSupplyUnit - sets gaming powerSupplyUnit for gaming computer
func (b *gamingComputerBuilder) setPowerSupplyUnit() {
	b.powerSypplyUnit = "Some gaming powerSupplyUnit"
}

// setRAM - sets gaming RAM for gaming computer
func (b *gamingComputerBuilder) setRAM() {
	b.ram = "Some gaming RAM"
}

// setFAN - sets gaming FAN for gaming computer
func (b *gamingComputerBuilder) setFAN() {
	b.fan = "Some gaming FAN"
}

// getComputer - returns gaming computer
func (b *gamingComputerBuilder) getComputer() *Computer {
	return &Computer{
		motherBoard:     b.motherBoard,
		systemBlock:     b.systemBlock,
		cpu:             b.cpu,
		gpu:             b.gpu,
		powerSypplyUnit: b.powerSypplyUnit,
		ram:             b.ram,
		fan:             b.fan,
	}
}

// director - for building computer
type director struct {
	computerBuilder ComputerBuilder
}

// newDirector - creates new Director instance
func newDirector(c ComputerBuilder) *director {
	return &director{
		computerBuilder: c,
	}
}

// setBuilder - sets new builder type
func (d *director) setBuilder(b ComputerBuilder) {
	d.computerBuilder = b
}

// buildComputer - returns new computer
func (d *director) buildComputer() Computer {
	d.computerBuilder.setMotherBoard()
	d.computerBuilder.setSystemBlock()
	d.computerBuilder.setCPU()
	d.computerBuilder.setGPU()
	d.computerBuilder.setPowerSupplyUnit()
	d.computerBuilder.setRAM()
	d.computerBuilder.setFAN()
	return *d.computerBuilder.getComputer()
}
