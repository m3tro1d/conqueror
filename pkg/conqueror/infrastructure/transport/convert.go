package transport

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
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

func buildListTasksSpecification(request listTasksRequest) query.ListTasksSpecification {
	var sortSettings *query.TasksSortSettings
	if request.SortSettings != nil {
		sortSettings = &query.TasksSortSettings{
			Field: query.TasksSortField(request.SortSettings.Field),
			Order: query.SortOrder(request.SortSettings.Order),
		}
	}

	return query.ListTasksSpecification{
		Sort:          sortSettings,
		ShowCompleted: request.ShowCompleted,
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
