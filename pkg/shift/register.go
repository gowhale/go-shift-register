// Package shift contains code to control Shift Registers
package shift

import (
	"fmt"
	"log"
	"time"
)

type Register interface {
	Clear()
	PushBit()
	ShowCombo(combo []int) error
	NOutputs() int
}

// NewRegisterHardware returns a Register struct
// Requires RpioProcessor and pin numbers for serPin, srclkPin, rclkPin
// Also the amount of bits it controls i.e. if 2 8 bit shift registers are daisy chaned 16 bits
func NewRegisterHardware(tr RpioProcessor, serPin, srclkPin, rclkPin, bits int) Register {
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

	sr := RegisterHardware{srclk: srclk,
		ser:     ser,
		rclk:    rclk,
		Outputs: outputBits}

	sr.Clear()
	return &sr
}

// RegisterHardware represents a shift register
type RegisterHardware struct {
	srclk   PinProcessor
	ser     PinProcessor
	rclk    PinProcessor
	Outputs []int
}

// NOutputs returns amount of outputs
func (sr *RegisterHardware) NOutputs() int {
	return len(sr.Outputs)
}

// ShowCombo will send a bit combination to the outputs of register
// Note: Will always clear the display to start
func (sr *RegisterHardware) ShowCombo(combo []int) error {
	if len(sr.Outputs) != len(combo) {
		return fmt.Errorf("outputs not same length as bits")
	}
	sr.Outputs = combo
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
	return nil
}

// PushBit will push the bit on the ser pin to Q1
// Will push all bits to next Q
func (sr *RegisterHardware) PushBit() {
	highTime := time.Nanosecond
	sr.srclk.High()
	time.Sleep(highTime)
	sr.srclk.Low()
	sr.rclk.High()
	time.Sleep(highTime)
	sr.rclk.Low()
}

// Clear will set all Q outputs to 0
func (sr *RegisterHardware) Clear() {
	sr.ser.Low()
	for i := 0; i < len(sr.Outputs)+1; i++ {
		log.Printf("Clearing Q%d", i)
		sr.PushBit()
	}
}
