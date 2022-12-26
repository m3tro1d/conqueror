import React from 'react'
import styles from './TasksPanel.module.css'
import Task from "./Task/Task";

function TasksPanel() {
    // const [tasks, setTasks] = useState([])
    // useEffect(() => {
    //     tasksApi
    //         .listTasks()
    //         .then(response => setTasks(response.tasks))
    //         .catch(() => alert('YOU FUCKED UP!!!'))
    // }, [])

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
        <div className={styles.tasksPanel}>
            <ul className={styles.tasksList}>
                {tasks.map(task => (<Task key={task.id} task={task} />))}
            </ul>
        </div>
    )
}

export default TasksPanel
