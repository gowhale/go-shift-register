// Package main in Christmas lights puts on a pretty light show!
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
	sr := shift.NewRegisterHardware(tr, 5, 6, 13, 32)

	defer func() {
		sr.Clear()
		if err := tr.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err := merryChristmas(sr); err != nil {
		log.Fatalln(err)
	}
}

func everyOtherLight(bits, offset int) []int {
	lights := []int{}
	for i := 0; i < bits; i++ {
		if (i+offset)%2 == 0 {
			lights = append(lights, 0)
		} else {
			lights = append(lights, 1)
		}
	}
	return lights
}

func merryChristmas(sr shift.Register) error {
	for {
		for count := 0; count < sr.NOutputs(); count++ {
			combo := lightAddition(sr.NOutputs(), count)
			if err := sr.ShowCombo(combo); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 300)
		}

		combo := everyOtherLight(sr.NOutputs(), 0)
		if err := sr.ShowCombo(combo); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 200)
		combo = everyOtherLight(sr.NOutputs(), 1)
		if err := sr.ShowCombo(combo); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 200)
		// }
	}
}

func lightAddition(bits, offset int) []int {
	lights := []int{}
	for x := 0; x < bits; x++ {
		lights = append(lights, 0)
	}
	for x := 0; x < offset; x++ {
		lights[x] = 1
	}
	return lights
}
