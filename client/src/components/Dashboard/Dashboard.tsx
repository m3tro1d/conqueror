import React, { useEffect, useState } from 'react'
import Header from './Header/Header'
import { tasksApi } from '../../api/api'

function Dashboard(): JSX.Element {
    const [tasks, setTasks] = useState([])
    useEffect(() => {
        tasksApi
            .listTasks()
            .then(response => setTasks(response.tasks))
            .catch(() => alert('YOU FUCKED UP!!!'))
    }, [])

    return (
        <div>
            <Header />
            <ul>
                {tasks.map(task => (
                    <li>{task}</li>
                ))}
            </ul>
        </div>
    )
}

export default Dashboard
