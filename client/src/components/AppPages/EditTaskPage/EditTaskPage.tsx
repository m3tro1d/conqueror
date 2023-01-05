import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import TaskForm, { TaskData } from '../common/TaskForm/TaskForm'
import { tasksApi } from '../../../api/api'

function EditTaskPage() {
    const { id } = useParams()
    const [task, setTask] = useState()

    useEffect(() => {
        if (!id) {
            alert('Invalid task id')
            return
        }

        tasksApi
            .getTask(id)
            .then(response => setTask(response.task))
            .catch(() => alert('Failed to fetch task'))
    }, [])

    const onSubmit = async (task: TaskData) => {
        if (!id) {
            alert('Invalid task id')
            return
        }

        tasksApi
            .updateTask(id, {
                title: task.title,
                description: task.description,
                due_date: task.dueDate,
                subject_id: task.subjectId,
            })
            .then(() => window.location.assign('/tasks'))
            .catch(() => alert('Failed to update task.'))
    }

    return (
        <div>
            <TaskForm
                onSubmit={onSubmit}
                task={task}
            />
        </div>
    )
}

export default EditTaskPage
