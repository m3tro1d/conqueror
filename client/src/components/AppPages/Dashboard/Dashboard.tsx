import React from 'react'
import TasksPanel from './TasksPanel/TasksPanel'
import NotesPanel from './NotesPanel/NotesPanel'
import styles from './Dashboard.module.css'

function Dashboard(): JSX.Element {
    return (
        <div className={styles.dashboard}>
            <TasksPanel />
            <NotesPanel />
        </div>
    )
}

export default Dashboard
