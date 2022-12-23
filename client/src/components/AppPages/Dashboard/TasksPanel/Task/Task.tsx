import React from 'react'
import styles from './Task.module.css'

type Tag = {
    id: string
    name: string
}

type Task = {
    id: string
    due_date: string
    title: string
    description: string
    tags: Tag[]
    subject_id: string | null
}

type TaskProps = {
    task: Task
}

function Task({ task }: TaskProps) {
    return (
        <li className={styles.taskItem}>
            <span className={styles.checkbox}></span>
            {task.title}
        </li>
    )
}

export default Task
