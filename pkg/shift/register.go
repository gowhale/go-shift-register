package shift

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func NewShiftRegister(tr RpioProcessor, serPin, srclkPin, rclkPin, bits int) ShiftRegister {
	srclk := tr.Pin(srclkPin)
	srclk.Output()
	srclk.Low()

	ser := tr.Pin(serPin)
	ser.Output()
	ser.Low()

	rclk := tr.Pin(rclkPin)
	rclk.Output()
	rclk.Low()

	outputBits := []int{}
	for i := 0; i < bits; i++ {
		outputBits = append(outputBits, 0)
	}

	sr := ShiftRegister{srclk: srclk,
		ser:     ser,
		rclk:    rclk,
		outputs: outputBits}

	sr.Clear()
	return sr
}

type ShiftRegister struct {
	srclk   PinProcessor
	ser     PinProcessor
	rclk    PinProcessor
	outputs []int
}

func (sr *ShiftRegister) ShowCombo(combo []int) {
	sr.Clear()
	for i, j := 0, len(combo)-1; i < j; i, j = i+1, j-1 {
		combo[i], combo[j] = combo[j], combo[i]
	}
	for _, bit := range combo {
		if bit == 1 {
			sr.ser.High()
		} else {
			sr.ser.Low()
		}
		sr.PushBit()
	}
	sr.PushBit()
}

func (sr *ShiftRegister) OnOff() {
	for count := 0; count < 50; count++ {
		sr.ser.Low()
		log.Println("OFF")
		sr.PushBit()
		time.Sleep(time.Second)
		sr.ser.High()
		log.Println("ON")
		sr.PushBit()
		time.Sleep(time.Second)
	}
}

func (sr *ShiftRegister) PushBit() {
	highTime := time.Millisecond
	sr.srclk.High()
	time.Sleep(highTime)
	sr.srclk.Low()
	sr.rclk.High()
	time.Sleep(highTime)
	sr.rclk.Low()
}

func (sr *ShiftRegister) Clear() {
	sr.ser.Low()
	for i := 0; i < len(sr.outputs)+1; i++ {
		log.Printf("Clearing Q%d", i)
		sr.PushBit()
	}
}

type RpioProc struct{}

type TermRPIO struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name rpioProcessor -inpkg --filename rpio_processor_mock.go
type RpioProcessor interface {
	Open() (err error)
	Close() (err error)
	Pin(p int) PinProcessor
}

func (*RpioProc) Open() (err error) {
	return rpio.Open()
}

func (*RpioProc) Close() (err error) {
	return rpio.Close()
}

func (*RpioProc) Pin(p int) PinProcessor {
	return rpio.Pin(p)
}

func (*TermRPIO) Open() (err error) {
	log.Println("Opening")
	return nil
}

func (*TermRPIO) Close() (err error) {
	log.Println("Closing")
	return nil
}

func (*TermRPIO) Pin(p int) PinProcessor {
	return &TermPin{
		bcm: p,
	}
}

type TermPin struct {
	bcm int
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name pinProcessor -inpkg --filename pin_processor_mock.go
type PinProcessor interface {
	Output()
	Low()
	High()
}

func (t *TermPin) Output() {
	log.Printf("pin=%d mode=OUTPUT", t.bcm)
}

func (t *TermPin) Low() {
	log.Printf("pin=%d val=LOW", t.bcm)
}

func (t *TermPin) High() {
	log.Printf("pin=%d val=HIGH", t.bcm)
}
