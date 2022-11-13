package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/app/service"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
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

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) (DependencyContainer, error) {
	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewUserRepository(ctx, conn)
	userService := service.NewUserService(userRepository)

	subjectRepository := mysql.NewSubjectRepository(ctx, conn)
	subjectService := service.NewSubjectService(subjectRepository, userRepository)

	taskRepository := mysql.NewTaskRepository(ctx, conn)
	taskService := service.NewTaskService(taskRepository, userRepository)

	taskTagRepository := mysql.NewTaskTagRepository(ctx, conn)
	taskTagService := service.NewTaskTagService(taskTagRepository, userRepository)

	noteRepository := mysql.NewNoteRepository(ctx, conn)
	noteService := service.NewNoteService(noteRepository, userRepository)

	noteTagRepository := mysql.NewNoteTagRepository(ctx, conn)
	noteTagService := service.NewNoteTagService(noteTagRepository, userRepository)

	userQueryService := mysql.NewUserQueryService(conn)
	taskQueryService := mysql.NewTaskQueryService(conn)
	noteQueryService := mysql.NewNoteQueryService(conn)

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
