import React from 'react'
import styles from './TasksPage.module.css'
import Task from '../common/Task/Task'

function TasksPage() {
    const tasks = [
        {
            id: '1',
            due_date: '2006-01-02 15:04:05.999999999 -0700 MST',
            title: 'Listen up',
            description: 'Here we go',
            tags: [
                {
                    id: 't1',
                    name: 'huinya',
                },
                {
                    id: 't2',
                    name: 'study',
                },
            ],
            subject_id: null,
        },
        {
            id: '2',
            due_date: '2006-01-03 15:04:05.999999999 -0700 MST',
            title: 'Its labour day and my grandpa just ate 7 fucking hot dogs',
            description: 'Break stuff',
            tags: [
                {
                    id: 't1',
                    name: 'work',
                },
            ],
            subject_id: 's1',
        },
        {
            id: '3',
            due_date: '2006-01-03 15:04:05.999999999 -0700 MST',
            title: 'But I forgot my pen',
            description: 'Shit the bed again',
            tags: [],
            subject_id: 's2',
        },
    ]

    return (
        <div className={styles.tasksPage}>
            <ul className={styles.tasksList}>
                {tasks.map(task => (<Task key={task.id} task={task} />))}
            </ul>
        </div>
    )
}

export default TasksPage
