package mysql

import "time"

type sqlxUser struct {
	ID       binaryUUID     `db:"id"`
	Login    string         `db:"login"`
	Password string         `db:"password"`
	AvatarID nullBinaryUUID `db:"avatar_id"`
}

type sqlxImage struct {
	ID   binaryUUID `db:"id"`
	Path string     `db:"path"`
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
	Status      int            `db:"status"`
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

type sqlxQueryUser struct {
	ID         binaryUUID     `db:"id"`
	Login      string         `db:"login"`
	Password   string         `db:"password"`
	AvatarID   nullBinaryUUID `db:"avatar_id"`
	AvatarPath *string        `db:"avatar_path"`
}

type sqlxQuerySubject struct {
	ID    binaryUUID `db:"id"`
	Title string     `db:"title"`
}

type sqlxNoteTag struct {
	ID     binaryUUID `db:"id"`
	Name   string     `db:"name"`
	UserID binaryUUID `db:"user_id"`
}

type sqlxQueryTask struct {
	ID           binaryUUID     `db:"id"`
	DueDate      time.Time      `db:"due_date"`
	Title        string         `db:"title"`
	Description  string         `db:"description"`
	Status       int            `db:"status"`
	SubjectID    nullBinaryUUID `db:"subject_id"`
	SubjectTitle *string        `db:"subject_title"`
}

type sqlxQueryNote struct {
	ID           binaryUUID     `db:"id"`
	Title        string         `db:"title"`
	Content      string         `db:"content"`
	UpdatedAt    time.Time      `db:"updated_at"`
	SubjectID    nullBinaryUUID `db:"subject_id"`
	SubjectTitle *string        `db:"subject_title"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
