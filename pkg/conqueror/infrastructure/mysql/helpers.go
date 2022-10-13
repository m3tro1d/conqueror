package mysql

type sqlxUser struct {
	ID       binaryUUID
	Login    string
	Password string
	Nickname string
}

type migrationInfo struct {
	Version  int
	FilePath string
}
