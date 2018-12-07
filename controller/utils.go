package controller

type State int

const (
	NoRegister State = iota // value --> 0
	OK      				// value --> 1
	WrongParams				// value --> 2
)
