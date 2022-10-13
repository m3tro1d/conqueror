package mysql

type sqlxUser struct {
	ID       binaryUUID `db:"id"`
	Login    string     `db:"login"`
	Password string     `db:"password"`
	Nickname string     `db:"nickname"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
