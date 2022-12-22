import React from 'react'
import TimetablePanel from './TimetablePanel/TimetablePanel'
import TasksPanel from './TasksPanel/TasksPanel'
import NotesPanel from './NotesPanel/NotesPanel'

function Dashboard(): JSX.Element {
    return (
        <div>
            <TimetablePanel />
            <div>
                <TasksPanel />
                <NotesPanel />
            </div>
        </div>
    )
}

export default Dashboard
