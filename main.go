// Package main runs the shopping list
package main

import (
	"flag"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func NewShiftRegister(tr rpioProcessor, serPin, srclkPin, rclkPin, bits int) ShiftRegister {
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
	srclk   pinProcessor
	ser     pinProcessor
	rclk    pinProcessor
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

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	flag.Parse()

	var tr rpioProcessor
	tr = &termRPIO{}
	if !*debugMode {
		tr = &rpioProc{}
	}

	if err := tr.Open(); err != nil {
		log.Fatalln(err)
	}
	sr := NewShiftRegister(tr, 16, 22, 27, 8)
	sr2 := NewShiftRegister(tr, 21, 20, 12, 8)
	sr3 := NewShiftRegister(tr, 5, 6, 13, 16)

	defer func() {
		sr.Clear()
		sr2.Clear()
		sr3.Clear()
		if err := tr.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	sr.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1})
	sr2.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1})
	sr3.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0})
	time.Sleep(time.Second * 5)
	sr.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0})
	sr2.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0})
	sr3.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0})
	time.Sleep(time.Second * 5)
	sr.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1})
	sr2.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1})
	sr3.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1})
	time.Sleep(time.Second * 5)
	sr.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1})
	sr2.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1})
	sr3.ShowCombo([]int{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0})
	time.Sleep(time.Second * 5)
}

type rpioProc struct{}

type termRPIO struct{}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name rpioProcessor -inpkg --filename rpio_processor_mock.go
type rpioProcessor interface {
	Open() (err error)
	Close() (err error)
	Pin(p int) pinProcessor
}

func (*rpioProc) Open() (err error) {
	return rpio.Open()
}

func (*rpioProc) Close() (err error) {
	return rpio.Close()
}

func (*rpioProc) Pin(p int) pinProcessor {
	return rpio.Pin(p)
}

func (*termRPIO) Open() (err error) {
	log.Println("Opening")
	return nil
}

func (*termRPIO) Close() (err error) {
	log.Println("Closing")
	return nil
}

func (*termRPIO) Pin(p int) pinProcessor {
	return &termPin{
		bcm: p,
	}
}

type termPin struct {
	bcm int
}

//go:generate go run github.com/vektra/mockery/cmd/mockery -name pinProcessor -inpkg --filename pin_processor_mock.go
type pinProcessor interface {
	Output()
	Low()
	High()
}

func (t *termPin) Output() {
	log.Printf("pin=%d mode=OUTPUT", t.bcm)
}

func (t *termPin) Low() {
	log.Printf("pin=%d val=LOW", t.bcm)
}

func (t *termPin) High() {
	log.Printf("pin=%d val=HIGH", t.bcm)
}
