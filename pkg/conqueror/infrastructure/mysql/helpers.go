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
	ID          binaryUUID     `db:"id"`
	UserID      binaryUUID     `db:"user_id"`
	DueDate     time.Time      `db:"due_date"`
	Title       string         `db:"title"`
	Description string         `db:"description"`
	SubjectID   nullBinaryUUID `db:"subject_id"`
}

type sqlxTaskTag struct {
	ID     binaryUUID `db:"id"`
	Name   string     `db:"name"`
	UserID binaryUUID `db:"user_id"`
}

type sqlxNote struct {
	ID        binaryUUID     `db:"id"`
	UserID    binaryUUID     `db:"user_id"`
	Title     string         `db:"title"`
	Content   string         `db:"content"`
	UpdatedAt time.Time      `db:"updated_at"`
	SubjectID nullBinaryUUID `db:"subject_id"`
}

type sqlxNoteTag struct {
	ID     binaryUUID `db:"id"`
	Name   string     `db:"name"`
	UserID binaryUUID `db:"user_id"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
