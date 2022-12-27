import React, { useEffect, useState } from 'react'
import styles from './TasksPanel.module.css'
import Task from '../../common/Task/Task'
import { tasksApi } from '../../../../api/api'

function TasksPanel() {
    const [tasks, setTasks] = useState([])

    const updateTasks = () => {
        tasksApi
            .listTasks()
            .then(response => setTasks(response.tasks))
            .catch(() => alert('Failed to fetch tasks.'))
    }

    useEffect(updateTasks, [])

    return (
        <div className={styles.tasksPanel}>
            <ul className={styles.tasksList}>
                {tasks.map(task => (
                    <Task key={task['id']} task={task} />
                ))}
            </ul>
        </div>
    )
}

export default TasksPanel
