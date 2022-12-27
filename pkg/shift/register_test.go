package shift

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type registerSuite struct {
	suite.Suite

	testRegister RegisterHardware
}

func (t *registerSuite) SetupTest() {
	t.testRegister = RegisterHardware{
		&termPin{}, &termPin{}, &termPin{}, []int{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func TestGuiSuite(t *testing.T) {
	suite.Run(t, new(registerSuite))
}

func (t *registerSuite) Test_NewRegister() {
	r := NewRegisterHardware(&TermRPIO{}, 1, 2, 3, 8)
	t.Equal(8, r.NOutputs())
}

func (t *registerSuite) Test_ShowCombo_Pass() {
	testCombo := []int{0, 1, 0, 1, 0, 1, 0, 1}
	err := t.testRegister.ShowCombo(testCombo)
	t.Equal(testCombo, t.testRegister.Outputs)
	t.Nil(err)
}

func (t *registerSuite) Test_ShowCombo_Fail() {
	testCombo := []int{0, 1, 0, 1, 0, 1, 0}
	err := t.testRegister.ShowCombo(testCombo)
	t.EqualError(err, "outputs not same length as bits")
}
