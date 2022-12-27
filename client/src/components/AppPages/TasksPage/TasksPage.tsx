import React, { useEffect, useState } from 'react'
import styles from './TasksPage.module.css'
import Task from '../common/Task/Task'
import AddTaskForm from '../Dashboard/TasksPanel/AddTaskForm/AddTaskForm'
import { tasksApi } from '../../../api/api'

function TasksPage() {
    const [tasks, setTasks] = useState([])

    const updateTasks = () => {
        tasksApi
            .listTasks()
            .then(response => setTasks(response.tasks))
            .catch(() => alert('Failed to fetch tasks.'))
    }

    useEffect(updateTasks, [])

    return (
        <div className={styles.tasksPage}>
            <AddTaskForm />
            <ul className={styles.tasksList}>
                {tasks.map(task => (
                    <Task key={task['id']} task={task} />
                ))}
            </ul>
        </div>
    )
}

export default TasksPage
