package State1

import "fmt"

type OnFailImp struct {
	m *StateMachine
}

func (o OnFailImp) OnCreate() error {
	//TODO implement me
	panic("implement me")
}

func (o OnFailImp) OnSuccess() error {
	//TODO implement me
	panic("implement me")
}

func (o OnFailImp) OnFail() error {
	fmt.Println("OnFail")
	return nil
}

var _ State = (*OnFailImp)(nil)
