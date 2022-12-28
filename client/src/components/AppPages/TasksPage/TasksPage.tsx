import React from 'react'
import styles from './TasksPage.module.css'
import Task from '../common/Task/Task'
import AddTaskForm from '../Dashboard/TasksPanel/AddTaskForm/AddTaskForm'
import useTasks from '../../../hooks/useTasks'

function TasksPage() {
    const { tasks, updateTasks, removeTask } = useTasks()

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
                        removeTask={removeTask}
                    />
                ))}
            </ul>
        </div>
    )
}

export default TasksPage
