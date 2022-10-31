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
}

type changeTaskTitleRequest struct {
	NewTitle string `json:"new_title"`
}

type changeTaskDescriptionRequest struct {
	NewDescription string `json:"new_description"`
}
