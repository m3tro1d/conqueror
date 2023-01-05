import React, { useEffect, useState } from 'react'
import styles from './TasksPage.module.css'
import TaskForm, { TaskData } from '../common/TaskForm/TaskForm'
import useTasks from '../../../hooks/useTasks'
import Task from '../common/Task/Task'
import useDebounce from '../../../hooks/useDebounce'
import { tasksApi } from '../../../api/api'

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

    const onSubmit = async (task: TaskData) => {
        await tasksApi.createTask({
            due_date: task.dueDate,
            title: task.title,
            description: task.description,
            subject_id: task.subjectId,
        })
        updateTasks(debouncedQuery)
    }

    const [sort, setSort] = useState('status')
    useEffect(() => updateTasks(debouncedQuery, sort), [sort])

    return (
        <div className={styles.tasksPage}>
            <TaskForm onSubmit={onSubmit} />

            <label htmlFor="search" className={styles.label}>Search</label>
            <input
                type="text"
                name="search"
                className={styles.input}
                onChange={e => setQuery(e.target.value)}
            />

            <label htmlFor="sort" className={styles.label}>Sort by</label>
            <select
                name="sort"
                className={styles.input}
                onChange={e => setSort(e.target.value)}
            >
                <option value="status" selected>Status</option>
                <option value="title">Title</option>
            </select>

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
