import React from 'react'
import styles from './TasksPanel.module.css'
import Task from '../../common/Task/Task'
import useTasks from '../../../../hooks/useTasks'

function TasksPanel() {
    const {tasks, changeTaskStatus} = useTasks({
        showCompleted: false,
    })

    return (
        <div className={styles.tasksPanel}>
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
                    />
                ))}
            </ul>
        </div>
    )
}

export default TasksPanel
