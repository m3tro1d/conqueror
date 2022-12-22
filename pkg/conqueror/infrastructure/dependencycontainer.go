package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/app/service"
	"conqueror/pkg/conqueror/infrastructure/mysql"
)

type DependencyContainer interface {
	UserService() service.UserService
	SubjectService() service.SubjectService

	TaskService() service.TaskService
	TaskTagService() service.TaskTagService

	NoteService() service.NoteService
	NoteTagService() service.NoteTagService

	UserQueryService() query.UserQueryService
	TaskQueryService() query.TaskQueryService
	NoteQueryService() query.NoteQueryService
}

func NewDependencyContainer(ctx context.Context, db mysql.ClientContext) (DependencyContainer, error) {
	userRepository := mysql.NewUserRepository(ctx, db)
	userService := service.NewUserService(userRepository)

	subjectRepository := mysql.NewSubjectRepository(ctx, db)
	subjectService := service.NewSubjectService(subjectRepository, userRepository)

	taskRepository := mysql.NewTaskRepository(ctx, db)
	taskService := service.NewTaskService(taskRepository, userRepository)

	taskTagRepository := mysql.NewTaskTagRepository(ctx, db)
	taskTagService := service.NewTaskTagService(taskTagRepository, userRepository)

	noteRepository := mysql.NewNoteRepository(ctx, db)
	noteService := service.NewNoteService(noteRepository, userRepository)

	noteTagRepository := mysql.NewNoteTagRepository(ctx, db)
	noteTagService := service.NewNoteTagService(noteTagRepository, userRepository)

	userQueryService := mysql.NewUserQueryService(db)
	taskQueryService := mysql.NewTaskQueryService(db)
	noteQueryService := mysql.NewNoteQueryService(db)

	return &dependencyContainer{
		userService:    userService,
		subjectService: subjectService,

		taskService:    taskService,
		taskTagService: taskTagService,

		noteService:    noteService,
		noteTagService: noteTagService,

		userQueryService: userQueryService,
		taskQueryService: taskQueryService,
		noteQueryService: noteQueryService,
	}, nil
}

type dependencyContainer struct {
	userService    service.UserService
	subjectService service.SubjectService

	taskService    service.TaskService
	taskTagService service.TaskTagService

	noteService    service.NoteService
	noteTagService service.NoteTagService

	userQueryService query.UserQueryService
	taskQueryService query.TaskQueryService
	noteQueryService query.NoteQueryService
}

func (container *dependencyContainer) UserService() service.UserService {
	return container.userService
}

func (container *dependencyContainer) SubjectService() service.SubjectService {
	return container.subjectService
}

func (container *dependencyContainer) TaskService() service.TaskService {
	return container.taskService
}

func (container *dependencyContainer) TaskTagService() service.TaskTagService {
	return container.taskTagService
}

func (container *dependencyContainer) NoteService() service.NoteService {
	return container.noteService
}

func (container *dependencyContainer) NoteTagService() service.NoteTagService {
	return container.noteTagService
}

func (container *dependencyContainer) UserQueryService() query.UserQueryService {
	return container.userQueryService
}

func (container *dependencyContainer) TaskQueryService() query.TaskQueryService {
	return container.taskQueryService
}

func (container *dependencyContainer) NoteQueryService() query.NoteQueryService {
	return container.noteQueryService
}
