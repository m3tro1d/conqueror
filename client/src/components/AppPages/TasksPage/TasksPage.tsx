import React from 'react'
import styles from './TasksPage.module.css'
import AddTaskForm from './AddTaskForm/AddTaskForm'
import useTasks from '../../../hooks/useTasks'
import Task from '../common/Task/Task'

function TasksPage() {
    const { tasks, updateTasks, changeTaskStatus, removeTask } = useTasks({
        showCompleted: true,
        sortField: 'status',
        sortOrder: 'asc',
    })

    return (
        <div className={styles.tasksPage}>
            <AddTaskForm updateTasks={updateTasks} />

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
                    />
                ))}
            </ul>
        </div>
    )
}

export default TasksPage
