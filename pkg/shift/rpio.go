package shift

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

// RpioProc is used when connected to the rapsberry Pi
type RpioProc struct{}

// TermRPIO is used to print functionality to terminal
type TermRPIO struct{}

// RpioProcessor is an interface which holds GPIO functionality
//go:generate go run github.com/vektra/mockery/cmd/mockery -name rpioProcessor -inpkg --filename rpio_processor_mock.go
type RpioProcessor interface {
	Open() (err error)
	Close() (err error)
	Pin(p int) PinProcessor
}

// Open and memory map GPIO memory range from /dev/mem . Some reflection magic is used to convert it to a unsafe []uint32 pointer
func (*RpioProc) Open() (err error) {
	return rpio.Open()
}

// Close unmaps GPIO memory
func (*RpioProc) Close() (err error) {
	return rpio.Close()
}

// Pin returns a rpio Pin
func (*RpioProc) Pin(p int) PinProcessor {
	return rpio.Pin(p)
}

// Open emulates RPIO Open functionality
func (*TermRPIO) Open() (err error) {
	log.Println("Opening")
	return nil
}

// Close emulates RPIO Close functionality
func (*TermRPIO) Close() (err error) {
	log.Println("Closing")
	return nil
}

// Pin returns a TermPin struct to show the pin being turned on and off
func (*TermRPIO) Pin(p int) PinProcessor {
	return &termPin{
		bcm: p,
	}
}

type termPin struct {
	bcm int
}

// PinProcessor is an interface which has Pin functionality
//go:generate go run github.com/vektra/mockery/cmd/mockery -name pinProcessor -inpkg --filename pin_processor_mock.go
type PinProcessor interface {
	Output()
	Low()
	High()
}

// Output mocking that a pin has been set to output
func (t *termPin) Output() {
	log.Printf("pin=%d mode=OUTPUT", t.bcm)
}

// Low mocks that a pin's output has been set to 0
func (t *termPin) Low() {
	log.Printf("pin=%d val=LOW", t.bcm)
}

// High mocks that a pin's output has been set to 1
func (t *termPin) High() {
	log.Printf("pin=%d val=HIGH", t.bcm)
}
