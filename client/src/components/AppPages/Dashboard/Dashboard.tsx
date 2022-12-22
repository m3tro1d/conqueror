import React from 'react'
import TimetablePanel from './TimetablePanel/TimetablePanel'
import TasksPanel from './TasksPanel/TasksPanel'
import NotesPanel from './NotesPanel/NotesPanel'
import styles from './Dashboard.module.css'

function Dashboard(): JSX.Element {
    return (
        <div className={styles.dashboard}>
            <TimetablePanel />
            <div className={styles.rightSidePanels}>
                <TasksPanel />
                <NotesPanel />
            </div>
        </div>
    )
}

export default Dashboard
