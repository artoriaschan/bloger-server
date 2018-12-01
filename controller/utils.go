package controller

type State int

const (
	OK State = iota // value --> 0
	NoRegister      // value --> 1
	WrongPassword   // value --> 2
)
