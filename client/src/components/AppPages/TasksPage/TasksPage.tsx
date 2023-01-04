import React, { useEffect, useState } from 'react'
import styles from './TasksPage.module.css'
import TaskForm from '../common/TaskForm/TaskForm'
import useTasks from '../../../hooks/useTasks'
import Task from '../common/Task/Task'
import useDebounce from '../../../hooks/useDebounce'

function TasksPage() {
    const { tasks, updateTasks, changeTaskStatus, removeTask } = useTasks({
        showCompleted: true,
        sortField: 'status',
        sortOrder: 'asc',
        for_today: false,
    })

    const [query, setQuery] = useState('')
    const debouncedQuery = useDebounce(query, 500)
    useEffect(() => updateTasks(debouncedQuery), [debouncedQuery])

    return (
        <div className={styles.tasksPage}>
            <TaskForm updateTasks={updateTasks} />

            <label htmlFor="search" className={styles.searchLabel}>Search</label>
            <input
                type="text"
                name="search"
                className={styles.searchBar}
                onChange={e => setQuery(e.target.value)}
            />

            <ul className={styles.tasksList}>
                {
                    tasks.length === 0 &&
                    <div className={styles.noTasksNotice}>No tasks</div>
                }
                {tasks.map(task => (
                    <Task
                        key={task['id']}
                        task={task}
                        changeTaskStatus={changeTaskStatus}
                        removeTask={removeTask}
                        showDate={true}
                    />
                ))}
            </ul>
        </div>
    )
}

export default TasksPage
