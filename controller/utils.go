package controller

type State int

const (
	NoRegister   State = 0
	OK           State = 1
	WrongParams  State = 2
	BadServer    State = 3
	BadDB        State = 4
	ExpiredToken State = 5
	NoToken      State = 6
)

const (
	JWTLogin      string = "100001"
	JWTAdminLogin string = "100002"
	JWTAutoLogin  string = "100003"
)
