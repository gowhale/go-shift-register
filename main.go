// Package main runs the shopping list
package main

import (
	"flag"
	"go-shift-register/pkg/shift"
	"log"
	"time"
)

func main() {
	var debugMode = flag.Bool("debug", false, "run in debug mode")
	flag.Parse()

	var tr shift.RpioProcessor
	tr = &shift.TermRPIO{}
	if !*debugMode {
		tr = &shift.RpioProc{}
	}

	if err := tr.Open(); err != nil {
		log.Fatalln(err)
	}
	sr := shift.NewShiftRegister(tr, 16, 22, 27, 8)
	sr2 := shift.NewShiftRegister(tr, 21, 20, 12, 8)
	sr3 := shift.NewShiftRegister(tr, 5, 6, 13, 16)

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
