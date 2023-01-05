package transport

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
	stderrors "errors"
	"github.com/gin-gonic/gin"
)

func queryUserToApi(user query.UserData) getUserResponse {
	var avatar *imageData
	if user.Avatar != nil {
		avatar = &imageData{
			ID:  user.Avatar.ImageID.String(),
			URL: user.Avatar.URL,
		}
	}

	return getUserResponse{
		UserID: user.UserID.String(),
		Login:  user.Login,
		Avatar: avatar,
	}
}

func querySubjectsToApi(subjects []query.SubjectData) []subjectData {
	result := make([]subjectData, 0, len(subjects))
	for _, subject := range subjects {
		result = append(result, subjectData{
			ID:    subject.ID.String(),
			Title: subject.Title,
		})
	}

	return result
}

func buildListTasksSpecification(ctx *gin.Context) (query.ListTasksSpecification, error) {
	showCompleted := true
	if ctx.Query("show_completed") == "false" {
		showCompleted = false
	}

	forToday := true
	if ctx.Query("for_today") == "false" {
		forToday = false
	}

	var sortSettings *query.TasksSortSettings
	field := ctx.Query("sort_field")
	order := ctx.Query("sort_order")
	if field != "" && order != "" {
		var queryField query.TasksSortField
		switch field {
		case "status":
			queryField = query.TasksSortFieldStatus
		case "title":
			queryField = query.TasksSortFieldTitle
		default:
			return query.ListTasksSpecification{}, stderrors.New("invalid field")
		}

		var queryOrder query.SortOrder
		switch order {
		case "asc":
			queryOrder = query.SortOrderAsc
		case "desc":
			queryOrder = query.SortOrderDesc
		default:
			return query.ListTasksSpecification{}, stderrors.New("invalid order")
		}

		sortSettings = &query.TasksSortSettings{
			Field: queryField,
			Order: queryOrder,
		}
	}

	return query.ListTasksSpecification{
		Query:         ctx.Query("query"),
		Sort:          sortSettings,
		ShowCompleted: showCompleted,
		ForToday:      forToday,
	}, nil
}

func buildListNotesSpecification(ctx *gin.Context) query.ListNotesSpecification {
	return query.ListNotesSpecification{
		Query: ctx.Query("query"),
	}
}

func queryTasksToApi(tasks []query.TaskData) []taskData {
	result := make([]taskData, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, queryTaskToApi(task))
	}

	return result
}

func queryTaskToApi(task query.TaskData) taskData {
	return taskData{
		ID:           task.ID.String(),
		DueDate:      task.DueDate,
		Title:        task.Title,
		Description:  task.Description,
		Status:       int(task.Status),
		SubjectID:    uuid.OptionalToString(task.SubjectID),
		SubjectTitle: task.SubjectTitle,
	}
}

func queryNotesToApi(notes []query.NoteData) []noteData {
	result := make([]noteData, 0, len(notes))
	for _, note := range notes {
		result = append(result, queryNoteToApi(note))
	}

	return result
}

func queryNoteToApi(note query.NoteData) noteData {
	return noteData{
		ID:        note.ID.String(),
		Title:     note.Title,
		Content:   note.Content,
		UpdatedAt: note.UpdatedAt.Unix(),
		SubjectID: uuid.OptionalToString(note.SubjectID),
	}
}
