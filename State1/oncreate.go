package State1

type OnCreateImp struct {
	m *StateMachine
}

func (o OnCreateImp) OnCreate() error {
	println("onCreate")

	//o.m.SetCurState(&OnSuccessImp{m: o.m})

	//放到service
	//o.m.SetCurState(&OnFailImp{m: o.m})
	return nil
}

func (o OnCreateImp) OnSuccess() error {
	//TODO implement me
	panic("implement me")
}

func (o OnCreateImp) OnFail() error {
	//TODO implement me
	panic("implement me")
}

var _ State = (*OnCreateImp)(nil)
