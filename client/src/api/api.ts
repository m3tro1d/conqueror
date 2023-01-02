import axios from 'axios'

type UserFormData = {
    login: string
    password: string
}

type SubjectData = {
    title: string
}

type TaskData = {
    due_date: Date
    title: string
    description: string
    subject_id?: string
}

type NoteData = {
    title: string
    content: string
    subject_id?: string
}

export type ListTasksSpecification = {
    showCompleted: boolean
    sortField?: string
    sortOrder?: string
}

function createInstance() {
    return axios.create({
        baseURL: 'http://localhost:8080/api/v1',
        headers: {
            'Access-Control-Allow-Origin': '*',
            'Access-Control-Allow-Methods': 'GET, POST, PATCH, DELETE',
            'X-Auth-Token': localStorage.getItem('token'),
        },
    })
}

const userApi = {
    login(data: UserFormData) {
        return createInstance()
            .post('/user/login', data)
            .then(response => response.data)
    },
    signup(data: UserFormData) {
        return createInstance()
            .post('/user', data)
    },
    changeAvatar(file: File) {
        const data = new FormData()
        data.set('avatar', file)

        return createInstance()
            .patch('/user/avatar', data)
    },
    getUser() {
        return createInstance()
            .get('/user')
            .then(response => response.data)
    },
}

const subjectApi = {
    listSubjects() {
        return createInstance()
            .get('/subjects')
            .then(response => response.data)
    },
    createSubject(data: SubjectData) {
        return createInstance()
            .post('/subject', data)
    },
    changeSubjectTitle(subjectId: string, title: string) {
        return createInstance()
            .patch(`/subject/${subjectId}/title`, {
                new_title: title,
            })
    },
    removeSubject(subjectId: string) {
        return createInstance()
            .delete(`/subject/${subjectId}`)
    },
}

const tasksApi = {
    listTasks(spec: ListTasksSpecification) {
        return createInstance()
            .get('/tasks', {
                params: {
                    show_completed: spec.showCompleted,
                    sort_field: spec.sortField,
                    sort_order: spec.sortOrder,
                },
            })
            .then(response => response.data)
    },
    createTask(data: TaskData) {
        return createInstance()
            .post('/task', data)
    },
    changeTaskTitle(taskId: string, title: string) {
        return createInstance()
            .patch(`/task/${taskId}/title`, {
                new_title: title,
            })
    },
    changeTaskDescription(taskId: string, description: string) {
        return createInstance()
            .patch(`/task/${taskId}/description`, {
                new_description: description,
            })
    },
    changeTaskStatus(taskId: string, status: number) {
        return createInstance()
            .patch(`/task/${taskId}/status`, {
                new_status: status,
            })
    },
    removeTask(taskId: string) {
        return createInstance()
            .delete(`/task/${taskId}`)
    },
}

const notesApi = {
    listNotes() {
        return createInstance()
            .get('/notes')
            .then(response => response.data)
    },
    createNote(data: NoteData) {
        return createInstance()
            .post('/note', data)
    },
    removeNote(noteId: string) {
        return createInstance()
            .delete(`/note/${noteId}`)
    },
}

export {
    userApi,
    subjectApi,
    tasksApi,
    notesApi,
}
