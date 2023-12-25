package main

import (
	"fmt"
)

// State 接口定义了状态的行为
type State interface {
	Execute()
}

// StateMachine 接口定义了状态机的行为
type StateMachine interface {
	GetState() State
	SetState(s State)
	Execute()
}

// ConcreteStateA 是 State 的一个具体状态实现
type ConcreteStateA struct{}

func (s *ConcreteStateA) Execute() {
	fmt.Println("Executing state A")
}

// ConcreteStateB 是 State 的一个具体状态实现
type ConcreteStateB struct{}

func (s *ConcreteStateB) Execute() {
	fmt.Println("Executing state B")
}

// ConcreteStateMachine 是 StateMachine 的一个具体实现
type ConcreteStateMachine struct {
	currentState State
}

func (sm *ConcreteStateMachine) GetState() State {
	return sm.currentState
}

func (sm *ConcreteStateMachine) SetState(s State) {
	sm.currentState = s
}

func (sm *ConcreteStateMachine) Execute() {
	sm.currentState.Execute()
}

func main() {
	stateMachine := &ConcreteStateMachine{}

	stateA := &ConcreteStateA{}
	stateB := &ConcreteStateB{}

	// 设置状态 A 并执行
	stateMachine.SetState(stateA)
	stateMachine.Execute()

	// 设置状态 B 并执行
	stateMachine.SetState(stateB)
	stateMachine.Execute()
}
