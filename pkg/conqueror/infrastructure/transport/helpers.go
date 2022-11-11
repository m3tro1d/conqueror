package transport

import "time"

type registerUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
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

type changeTaskTagsRequest struct {
	Tags []string `json:"tags"`
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
