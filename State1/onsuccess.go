package State1

import "fmt"

type OnSuccessImp struct {
	m *StateMachine
}

func (o OnSuccessImp) OnCreate() error {
	//TODO implement me
	panic("implement me")
}

func (o OnSuccessImp) OnSuccess() error {
	fmt.Println("OnSuccess")
	return nil
}

func (o OnSuccessImp) OnFail() error {
	//TODO implement me
	panic("implement me")
}

var _ State = (*OnSuccessImp)(nil)
