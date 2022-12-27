package shift

import (
	"fmt"
	"log"
	"strings"
)

// NewRegisterVirtual returns a RegisterVirtual struct
// Used for testing patters without connecting to LEDS etc.
func NewRegisterVirtual(bits int) Register {
	outputBits := []int{}
	for i := 0; i < bits; i++ {
		outputBits = append(outputBits, 0)
	}

	sr := RegisterVirtual{
		outputs: outputBits,
	}

	sr.Clear()
	return &sr
}

// RegisterVirtual represents a shift register
type RegisterVirtual struct {
	nextBit int
	outputs []int
}

// NOutputs returns amount of outputs
func (sr *RegisterVirtual) NOutputs() int {
	return len(sr.outputs)
}

func (sr *RegisterVirtual) ShowOutputs() {
	// Table headers
	padding := ""
	for range sr.outputs {
		padding = padding + strings.Repeat("#", 6)
	}
	log.Println(padding)

	header := ""
	for i := range sr.outputs {
		header = header + fmt.Sprintf("%5s|", fmt.Sprintf("Q%d", i+1))
	}

	log.Println(header)
	content := ""
	for _, q := range sr.outputs {
		content = content + fmt.Sprintf("%5d|", q)
	}
	log.Println(content)

	log.Println(padding)
}

// ShowCombo will send a bit combination to the outputs of register
// Note: Will always clear the display to start
func (sr *RegisterVirtual) ShowCombo(combo []int) error {
	if sr.NOutputs() != len(combo) {
		return fmt.Errorf("outputs not same length as bits")
	}
	for i, j := 0, len(combo)-1; i < j; i, j = i+1, j-1 {
		combo[i], combo[j] = combo[j], combo[i]
	}
	for _, bit := range combo {
		if bit == 1 {
			sr.nextBit = 1
		} else {
			sr.nextBit = 0
		}
		sr.PushBit()
	}
	sr.ShowOutputs()
	return nil
}

// PushBit will push the bit on the ser pin to Q1
// Will push all bits to next Q
func (sr *RegisterVirtual) PushBit() {
	newOutputs := make([]int, sr.NOutputs())
	for i := 0; i < sr.NOutputs()-1; i++ {
		newOutputs[i+1] = sr.outputs[i]
	}
	newOutputs[0] = sr.nextBit
	sr.outputs = newOutputs
}

// Clear will set all Q outputs to 0
func (sr *RegisterVirtual) Clear() {
	for i := 0; i < sr.NOutputs(); i++ {
		sr.outputs[i] = 0
	}
}
