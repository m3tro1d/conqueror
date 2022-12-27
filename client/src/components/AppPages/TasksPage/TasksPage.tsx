import React from 'react'
import styles from './TasksPage.module.css'
import Task from '../common/Task/Task'

function TasksPage() {
    const tasks = [
        {
            id: '1',
            due_date: '2006-01-02 15:04:05.999999999 -0700 MST',
            title: 'Listen up',
            description: 'Just get the fuck up',
            tags: [
                {
                    id: 't1',
                    name: 'ood',
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
            title: 'Fucked up job with a fucked up paying',
            description: 'Just get the fuck up x2',
            tags: [
                {
                    id: 't1',
                    name: 'automata',
                },
            ],
            subject_id: 's1',
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
