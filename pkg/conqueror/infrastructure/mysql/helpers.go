package mysql

import "time"

type sqlxUser struct {
	ID       binaryUUID `db:"id"`
	Login    string     `db:"login"`
	Password string     `db:"password"`
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

type sqlxQueryTaskTagWithTask struct {
	ID     binaryUUID `db:"id"`
	TaskID binaryUUID `db:"task_id"`
	Name   string     `db:"name"`
}

type sqlxQueryTaskTag struct {
	ID   binaryUUID `db:"id"`
	Name string     `db:"name"`
}

type sqlxQueryNote struct {
	ID        binaryUUID     `db:"id"`
	Title     string         `db:"title"`
	Content   string         `db:"description"`
	UpdatedAt time.Time      `db:"updated_at"`
	SubjectID nullBinaryUUID `db:"subject_id"`
}

type sqlxQueryNoteTagWithNote struct {
	ID     binaryUUID `db:"id"`
	NoteID binaryUUID `db:"note_id"`
	Name   string     `db:"name"`
}

type sqlxQueryNoteTag struct {
	ID   binaryUUID `db:"id"`
	Name string     `db:"name"`
}

type migrationInfo struct {
	Version  int
	FilePath string
}
