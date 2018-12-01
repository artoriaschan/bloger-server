package controller

type State int

const (
	OK State = iota 	// value --> 0
	NoRegister      	// value --> 1
	WrongPassword   	// value --> 2
	WrongFormatEmail	// value --> 3
	EmptyEmail			// value --> 4
	ExsitedEmail		// value --> 5
	WrongFormatPassword	// value --> 6
	EmpeyPassword		// value --> 7
	WrongFormatUsername	// value --> 8
	EmptyUsername		// value --> 9
	WrongFormatMobile	// value --> 10
)
