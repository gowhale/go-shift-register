package shift

import (
	"log"
	"time"
)

type shiftRegister struct {
	clock   *int
	serial  int
	outputs [8]int
}

func NewShiftRegister() shiftRegister {
	return shiftRegister{}
}

func ClockFreq(freq float64, clock *int) {
	for {
		*clock = 1
		log.Println(*clock)
		time.Sleep(time.Second / time.Duration((freq / 2)))
		*clock = 0
		log.Println(*clock)
		time.Sleep(time.Second / time.Duration((freq / 2)))
	}
}

type Pin struct {
	bcm int
	val int
}

func (p *Pin) ClockPin(clock *int) {
	if p.val != *clock {
		p.val = *clock
		log.Printf("pin=%d val=%d", p.bcm, p.val)
	}
}
