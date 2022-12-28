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
	UserID string `json:"user_id"`
	Login  string `json:"login"`
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

type changeTaskTitleRequest struct {
	NewTitle string `json:"new_title"`
}

type changeTaskDescriptionRequest struct {
	NewDescription string `json:"new_description"`
}

type changeTaskStatusRequest struct {
	NewStatus int `json:"new_status"`
}

type changeTaskTagsRequest struct {
	Tags []string `json:"tags"`
}

type createTaskTagRequest struct {
	Name string `json:"name"`
}

type changeTaskTagNameRequest struct {
	NewName string `json:"new_name"`
}

type createNoteRequest struct {
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	SubjectID *string `json:"subject_id"`
}

type changeNoteTitleRequest struct {
	NewTitle string `json:"new_title"`
}

type changeNoteContentRequest struct {
	NewContent string `json:"new_content"`
}

type changeNoteTagsRequest struct {
	Tags []string `json:"tags"`
}

type createNoteTagRequest struct {
	Name string `json:"name"`
}

type changeNoteTagNameRequest struct {
	NewName string `json:"new_name"`
}

type listSubjectsResponse struct {
	Subjects []subjectData `json:"subjects"`
}

type listTasksResponse struct {
	Tasks []taskData `json:"tasks"`
}

type listTaskTagsResponse struct {
	Tags []taskTagData `json:"tags"`
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

type listNoteTagsResponse struct {
	Tags []noteTagData `json:"tags"`
}

type noteData struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Tags      []noteTagData `json:"tags"`
	UpdatedAt string        `json:"updated_at"`
	SubjectID *string       `json:"subject_id"`
}

type noteTagData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
