package transport

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
	"github.com/gin-gonic/gin"
)

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

func buildListTasksSpecification(ctx *gin.Context) query.ListTasksSpecification {
	var showCompleted bool
	if ctx.Query("show_completed") == "false" {
		showCompleted = true
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
			return query.ListTasksSpecification{}
		}

		var queryOrder query.SortOrder
		switch order {
		case "asc":
			queryOrder = query.SortOrderAsc
		case "desc":
			queryOrder = query.SortOrderDesc
		default:
			return query.ListTasksSpecification{}
		}

		sortSettings = &query.TasksSortSettings{
			Field: queryField,
			Order: queryOrder,
		}
	}

	return query.ListTasksSpecification{
		Sort:          sortSettings,
		ShowCompleted: showCompleted,
	}
}

func queryTasksToApi(tasks []query.TaskData) []taskData {
	result := make([]taskData, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, taskData{
			ID:          task.ID.String(),
			DueDate:     task.DueDate,
			Title:       task.Title,
			Description: task.Description,
			Tags:        queryTaskTagsToApi(task.Tags),
			SubjectID:   uuid.OptionalToString(task.SubjectID),
		})
	}

	return result
}

func queryTaskTagsToApi(tags []query.TaskTagData) []taskTagData {
	result := make([]taskTagData, 0, len(tags))
	for _, tag := range tags {
		result = append(result, taskTagData{
			ID:   tag.ID.String(),
			Name: tag.Name,
		})
	}

	return result
}

func queryNotesToApi(notes []query.NoteData) []noteData {
	result := make([]noteData, 0, len(notes))
	for _, note := range notes {
		result = append(result, noteData{
			ID:        note.ID.String(),
			Title:     note.Title,
			Content:   note.Content,
			Tags:      queryNoteTagsToApi(note.Tags),
			UpdatedAt: note.UpdatedAt.String(),
			SubjectID: uuid.OptionalToString(note.SubjectID),
		})
	}

	return result
}

func queryNoteTagsToApi(tags []query.NoteTagData) []noteTagData {
	result := make([]noteTagData, 0, len(tags))
	for _, tag := range tags {
		result = append(result, noteTagData{
			ID:   tag.ID.String(),
			Name: tag.Name,
		})
	}

	return result
}
