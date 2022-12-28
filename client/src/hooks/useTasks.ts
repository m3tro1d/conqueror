import {useEffect, useState} from 'react'
import {ListTasksSpecification, tasksApi} from '../api/api'

function useTasks(spec: ListTasksSpecification) {
    const [tasks, setTasks] = useState([])

    const updateTasks = () => {
        tasksApi
            .listTasks(spec)
            .then(response => setTasks(response.tasks))
            .catch(() => alert('Failed to fetch tasks.'))
    }
    const changeTaskStatus = (id: string, status: number) => {
        tasksApi
            .changeTaskStatus(id, status)
            .then(updateTasks)
            .catch(() => alert('Failed to change task status.'))
    }
    const removeTask = (id: string) => {
        tasksApi
            .removeTask(id)
            .then(updateTasks)
            .catch(() => alert('Failed to remove task.'))
    }

    useEffect(updateTasks, [])

    return {
        tasks,
        updateTasks,
        changeTaskStatus,
        removeTask,
    }
}

export default useTasks
