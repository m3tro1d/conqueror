package transport

import (
	"conqueror/pkg/common/uuid"
	"conqueror/pkg/conqueror/app/query"
)

func queryTasksToApi(tasks []query.TaskData) []taskData {
	result := make([]taskData, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, taskData{
			ID:          task.ID.String(),
			DueDate:     task.DueDate.String(),
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

func queryNotesToApi(tasks []query.NoteData) []noteData {
	result := make([]noteData, 0, len(tasks))
	for _, task := range tasks {
		result = append(result, noteData{
			ID:        task.ID.String(),
			Title:     task.Title,
			Content:   task.Content,
			Tags:      queryNoteTagsToApi(task.Tags),
			SubjectID: uuid.OptionalToString(task.SubjectID),
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
