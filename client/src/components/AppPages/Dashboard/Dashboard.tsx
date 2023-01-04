import React from 'react'
import TasksPanel from './TasksPanel/TasksPanel'
import NotesPanel from './NotesPanel/NotesPanel'
import styles from './Dashboard.module.css'

function Dashboard(): JSX.Element {
    return (
        <div className={styles.dashboard}>
            <div className={styles.label}>
                <span>Tasks for today</span>
                <a href="/tasks">
                    <span className={'material-icons ' + styles.moreLink}>arrow_right_alt</span>
                </a>
            </div>
            <TasksPanel />

            <div className={styles.label}>
                <span>Recent notes</span>
                <a href="/notes">
                    <span className={'material-icons ' + styles.moreLink}>arrow_right_alt</span>
                </a>
            </div>
            <NotesPanel />
        </div>
    )
}

export default Dashboard
