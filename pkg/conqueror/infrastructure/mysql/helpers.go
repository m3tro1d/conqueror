package mysql

import "time"

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

type sqlxTask struct {
	ID          binaryUUID `db:"id"`
	UserID      binaryUUID `db:"user_id"`
	DueDate     time.Time  `db:"due_date"`
	Title       string     `db:"title"`
	Description string     `db:"description"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
