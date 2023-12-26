package State1

import "testing"

func mockService() bool {
	return true
}

func Test_NewMachine(t *testing.T) {

	machine := NewStateMachine()

	machine.GetCurState().OnCreate()

	res := mockService()

	if res {
		machine.SetCurState(&OnSuccessImp{m: machine})
		machine.GetCurState().OnSuccess()
	} else {
		machine.SetCurState(&OnFailImp{m: machine})
		machine.GetCurState().OnFail()
	}
	//service
	//文稿校对

	//machine.GetCurState().OnFail()
}
