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

	sr3 := shift.NewRegisterHardware(tr, 5, 6, 13, 32)

	defer func() {
		sr3.Clear()
		if err := tr.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := lightShow(sr3); err != nil {
		log.Fatalln(err)
	}
}

func lightShow(sr3 shift.Register) error {
	if err := sr3.ShowCombo([]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}); err != nil {
		return err
	}
	time.Sleep(time.Second * 5)

	if err := sr3.ShowCombo([]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0}); err != nil {
		return err
	}
	time.Sleep(time.Second * 5)

	if err := sr3.ShowCombo([]int{0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1}); err != nil {
		return err
	}
	time.Sleep(time.Second * 5)

	if err := sr3.ShowCombo([]int{1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0}); err != nil {
		return err
	}
	time.Sleep(time.Second * 5)
	return nil
}
