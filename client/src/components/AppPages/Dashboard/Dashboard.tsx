import React from 'react'
import Header from '../common/Header/Header'

function Dashboard(): JSX.Element {
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
            title: 'Get the fuck up',
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
            title: 'Get the fuck up x2',
            description: 'Just get the fuck up x2',
            tags: [
                {
                    id: 't1',
                    name: 'atomata',
                },
            ],
            subject_id: 's1',
        },
    ]

    return (
        <div>
            <Header />
            <ul>
                {tasks.map(task => (
                    <li
                        key={task.id}
                    >
                        {task.title}
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default Dashboard
