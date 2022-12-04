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
	sr := shift.NewRegister(tr, 16, 22, 27, 8)
	sr2 := shift.NewRegister(tr, 21, 20, 12, 8)
	sr3 := shift.NewRegister(tr, 5, 6, 13, 16)

	defer func() {
		sr.Clear()
		sr2.Clear()
		sr3.Clear()
		if err := tr.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := sr.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr2.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr3.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 5)
	if err := sr.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0}); err != nil {
		log.Fatalln(err)
	}
	if err := sr2.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0}); err != nil {
		log.Fatalln(err)
	}
	if err := sr3.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 5)
	if err := sr.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr2.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr3.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 5)
	if err := sr.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr2.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 1}); err != nil {
		log.Fatalln(err)
	}
	if err := sr3.ShowCombo([]int{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0}); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(time.Second * 5)
}
