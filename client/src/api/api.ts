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
    for_today: boolean
}

function createInstance() {
    return axios.create({
        baseURL: 'http://localhost/api/v1',
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
    listTasks(spec: ListTasksSpecification, query: string = '') {
        return createInstance()
            .get('/tasks', {
                params: {
                    show_completed: spec.showCompleted,
                    sort_field: spec.sortField,
                    sort_order: spec.sortOrder,
                    for_today: spec.for_today,
                    query: query,
                },
            })
            .then(response => response.data)
    },
    getTask(id: string) {
        return createInstance()
            .get(`/task/${id}`)
            .then(response => response.data)
    },
    createTask(data: TaskData) {
        return createInstance()
            .post('/task', data)
    },
    updateTask(taskId: string, data: TaskData) {
        return createInstance()
            .patch(`/task/${taskId}`, data)
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
    listNotes(query: string = '') {
        return createInstance()
            .get('/notes', {
                params: {
                    query: query,
                },
            })
            .then(response => response.data)
    },
    getNote(id: string) {
        return createInstance()
            .get(`/note/${id}`)
            .then(response => response.data)
    },
    createNote(data: NoteData) {
        return createInstance()
            .post('/note', data)
    },
    updateNote(id: string, data: NoteData) {
        return createInstance()
            .patch(`/note/${id}`, data)
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
