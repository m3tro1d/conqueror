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
    removeTask?: (id: string) => void
}

function Task({ task, removeTask }: TaskProps) {
    return (
        <li className={styles.taskItem}>
            <div className={styles.mainContent}>
                <span className={styles.checkbox}></span>
                <span>{task.title}</span>
            </div>

            <div className={styles.description}>
                {task.description}
            </div>

            <ul className={styles.tags}>
                {task.tags.map(tag => (
                    <li
                        key={tag.id}
                        className={styles.tag}
                    >
                        {tag.name}
                    </li>
                ))}
            </ul>

            {
                removeTask &&
                <a
                    href="#"
                    className={styles.removeButton}
                    onClick={e => {
                        e.preventDefault()
                        removeTask(task.id)
                    }}
                >
                    <span className="material-icons">delete</span>
                </a>
            }
        </li>
    )
}

export default Task
