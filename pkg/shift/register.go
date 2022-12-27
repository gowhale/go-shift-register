package shift

// Register is interface used for controlling virtual and hardware registers
type Register interface {
	Clear()
	PushBit()
	ShowCombo(combo []int) error
	NOutputs() int
}
