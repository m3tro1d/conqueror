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
    status: number
    tags: Tag[]
    subject_id: string | null
    subject_title: string | null
}

type TaskProps = {
    task: Task
    changeTaskStatus: (id: string, status: number) => void
    removeTask?: (id: string) => void
}

function Task({ task, changeTaskStatus, removeTask }: TaskProps) {
    const dueDate = new Date(Date.parse(task.due_date))
    const dateStr = `${dueDate.getFullYear()}-${dueDate.getMonth() + 1}-${dueDate.getDate()}`

    return (
        <li className={styles.taskItem}>
            <div className={styles.mainContent}>
                <span
                    className={styles.checkbox}
                    onClick={() => changeTaskStatus(task.id, 1 - task.status)}
                >
                    {
                        task.status === 1 &&
                        <span className={'material-icons ' + styles.mark}>done</span>
                    }
                </span>
                <span className={task.status === 1 ? styles.completed : ''}>{task.title}</span>
                <span className={styles.dueDate}>{dateStr}</span>
            </div>

            <div className={styles.description}>
                {task.description}
            </div>

            {
                task.subject_title !== null &&
                <span className={styles.subject}>
                    {task.subject_title}
                </span>
            }

            {
                task.tags.length !== 0 &&
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
            }

            {
                removeTask &&
                <div>
                    <span
                        className={'material-icons ' + styles.removeButton}
                        onClick={e => {
                            e.preventDefault()
                            removeTask(task.id)
                        }}
                    >
                        delete</span>
                </div>
            }
        </li>
    )
}

export default Task
