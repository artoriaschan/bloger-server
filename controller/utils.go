package controller

type State int

const (
	NoRegister State = 0
	OK State = 1
	WrongParams State = 2
	BadServer State = 3
	BadDB State = 4
	ExpireToken State = 5
)

const (
	JWTLogin string = "100001"
	JWTAdminLogin string = "100002"
)
