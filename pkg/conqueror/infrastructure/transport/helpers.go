package transport

import "time"

type registerUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

type getUserResponse struct {
	UserID string     `json:"user_id"`
	Login  string     `json:"login"`
	Avatar *imageData `json:"avatar"`
}

type createSubjectRequest struct {
	Title string `json:"title"`
}

type changeSubjectTitleRequest struct {
	NewTitle string `json:"new_title"`
}

type createTaskRequest struct {
	DueDate     time.Time `json:"due_date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SubjectID   *string   `json:"subject_id"`
}

type updateTaskRequest struct {
	NewTitle string `json:"new_title"`
}

type changeTaskStatusRequest struct {
	NewStatus int `json:"new_status"`
}

type createNoteRequest struct {
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	SubjectID *string `json:"subject_id"`
}

type updateNoteRequest struct {
	NewTitle string `json:"new_title"`
}

type listSubjectsResponse struct {
	Subjects []subjectData `json:"subjects"`
}

type listTasksResponse struct {
	Tasks []taskData `json:"tasks"`
}

type getTaskResponse struct {
	Task taskData `json:"task"`
}

type imageData struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type subjectData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type taskData struct {
	ID           string        `json:"id"`
	DueDate      time.Time     `json:"due_date"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Status       int           `json:"status"`
	Tags         []taskTagData `json:"tags"`
	SubjectID    *string       `json:"subject_id"`
	SubjectTitle *string       `json:"subject_title"`
}

type taskTagData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type listNotesResponse struct {
	Notes []noteData `json:"notes"`
}

type getNoteResponse struct {
	Note noteData `json:"note"`
}

type noteData struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Tags         []noteTagData `json:"tags"`
	UpdatedAt    int64         `json:"updated_at"`
	SubjectID    *string       `json:"subject_id"`
	SubjectTitle *string       `json:"subject_title"`
}

type noteTagData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
