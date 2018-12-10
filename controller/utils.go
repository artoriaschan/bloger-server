package controller

type State int

const (
	NoRegister     State = 0
	OK             State = 1
	WrongParams    State = 2
	BadServer      State = 3
	BadDB          State = 4
	NoRight        State = 5
	NoToken        State = 6
	FreezenAccount State = 7
	NoLogin        State = 8
)
