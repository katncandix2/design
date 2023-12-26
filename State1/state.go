package State1

type State interface {
	OnCreate() error
	OnSuccess() error
	OnFail() error
}
