package mysql

type sqlxUser struct {
	ID       uint
	Login    string
	Password string
	Nickname string
}
