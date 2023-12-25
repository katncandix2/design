package main

import (
	"fmt"
)

// State 类型代表状态机中的状态
type State string

// Event 类型代表触发状态转换的事件
type Event string

// StateMachine 结构体表示状态机
type StateMachine struct {
	state       State
	transitions map[State]map[Event]State
}

// NewStateMachine 创建并返回一个新的状态机实例
func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		state:       initialState,
		transitions: make(map[State]map[Event]State),
	}
}

// AddTransition 用于添加状态转换规则
func (sm *StateMachine) AddTransition(from State, event Event, to State) {
	if sm.transitions[from] == nil {
		sm.transitions[from] = make(map[Event]State)
	}
	sm.transitions[from][event] = to
}

// Transition 触发状态转换
func (sm *StateMachine) Transition(event Event) {
	if nextState, ok := sm.transitions[sm.state][event]; ok {
		sm.state = nextState
		fmt.Printf("Transitioned to %s\n", sm.state)
	} else {
		fmt.Printf("No transition for event %s from state %s\n", event, sm.state)
	}
}

// CurrentState 返回当前状态
func (sm *StateMachine) CurrentState() State {
	return sm.state
}

func main() {
	sm := NewStateMachine("idle")

	// 定义状态转换
	sm.AddTransition("idle", "start", "running")
	sm.AddTransition("running", "stop", "idle")
	sm.AddTransition("running", "pause", "paused")
	sm.AddTransition("paused", "resume", "running")

	// 测试状态转换
	sm.Transition("start")
	sm.Transition("pause")
	sm.Transition("resume")
	sm.Transition("stop")
}
