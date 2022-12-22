import axios from 'axios'

type UserFormData = {
    login: string
    password: string
}

type TaskData = {
    dueDate: Date
    title: string
    description: string
    subjectId?: string
}

type NoteData = {
    title: string
    content: string
    subjectId?: string
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

const authApi = {
    login(data: UserFormData) {
        return createInstance()
            .post('/user/login', data)
            .then(response => response.data)
    },
    signup(data: UserFormData) {
        return createInstance()
            .post('/user', data)
    },
}

const timetableApi = {
    createSubject() {

    },
    changeSubjectTitle() {
    },
    removeSubject() {

    },
}

const tasksApi = {
    listTasks() {
        return createInstance()
            .get('tasks')
            .then(response => response.data)
    },
    createTask(data: TaskData) {
        return createInstance()
            .post('task', data)
    },
    changeTaskTitle(taskId: string, title: string) {
        return createInstance()
            .patch(`task/${taskId}/title`, {
                new_title: title,
            })
    },
    changeTaskTags(taskId: string, tags: string[]) {
        return createInstance()
            .patch(`task/${taskId}/tags`, {
                tags: tags,
            })
    },
    changeTaskDescription(taskId: string, description: string) {
        return createInstance()
            .patch(`task/${taskId}/description`, {
                new_description: description,
            })
    },
    changeTaskStatus(taskId: string, status: number) {
        return createInstance()
            .patch(`task/${taskId}/status`, {
                new_status: status,
            })
    },
    removeTask(taskId: string) {
        return createInstance()
            .delete(`task/${taskId}`)
    },
}

const notesApi = {
    createNote(data: NoteData) {
        return createInstance()
            .post('note', data)
    },
}

export {
    authApi,
    timetableApi,
    tasksApi,
    notesApi,
}
