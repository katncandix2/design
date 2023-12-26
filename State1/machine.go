package State1

type StateMachine struct {
	CurState State
}

func NewStateMachine() *StateMachine {
	m := &StateMachine{}
	InitState := &OnCreateImp{m: m}
	m.SetCurState(InitState)
	return m
}

func (s *StateMachine) SetCurState(state State) {
	s.CurState = state
}

func (s *StateMachine) GetCurState() State {
	return s.CurState
}
