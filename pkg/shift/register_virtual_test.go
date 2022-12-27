package shift

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type virtualRegisterSuite struct {
	suite.Suite

	testRegister RegisterVirtual
}

func (t *virtualRegisterSuite) SetupTest() {
	t.testRegister = RegisterVirtual{
		0, []int{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func TestVirtualSuite(t *testing.T) {
	suite.Run(t, new(virtualRegisterSuite))
}

func (t *virtualRegisterSuite) Test_NewRegister() {
	r := NewRegisterVirtual(8)
	t.Equal(8, r.NOutputs())
}

func (t *virtualRegisterSuite) Test_ShowCombo_Pass() {
	testCombo := []int{0, 1, 0, 1, 0, 1, 0, 1}
	expectedOutputs := []int{0, 1, 0, 1, 0, 1, 0, 1}
	err := t.testRegister.ShowCombo(testCombo)
	t.Equal(expectedOutputs, t.testRegister.outputs)
	t.Nil(err)
}

func (t *virtualRegisterSuite) Test_ShowCombo_Fail() {
	testCombo := []int{0, 1, 0, 1, 0, 1, 0}
	err := t.testRegister.ShowCombo(testCombo)
	t.EqualError(err, "outputs not same length as bits")
}
