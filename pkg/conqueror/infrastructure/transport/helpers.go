package transport

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
