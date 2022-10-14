package mysql

type sqlxUser struct {
	ID       binaryUUID `db:"id"`
	Login    string     `db:"login"`
	Password string     `db:"password"`
	Nickname string     `db:"nickname"`
}

type sqlxSubject struct {
	ID     binaryUUID `db:"id"`
	UserID binaryUUID `db:"user_id"`
	Title  string     `db:"title"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
