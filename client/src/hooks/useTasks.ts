import { useEffect, useState } from 'react'
import { tasksApi } from '../api/api'

function useTasks() {
    const [tasks, setTasks] = useState([])

    const updateTasks = () => {
        tasksApi
            .listTasks()
            .then(response => setTasks(response.tasks))
            .catch(() => alert('Failed to fetch tasks.'))
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
        removeTask,
    }
}

export default useTasks
