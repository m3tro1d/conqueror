import React, { FormEvent, useState } from 'react'
import styles from './TaskForm.module.css'
import useSubjects from '../../../../hooks/useSubjects'

type Task = {
    id: string
    due_date: string
    title: string
    description: string
    status: number
    subject_id: string | null
    subject_title: string | null
}

export type TaskData = {
    dueDate: Date
    title: string
    description: string
    subjectId?: string
}

type TaskFormProps = {
    onSubmit?: (task: TaskData) => void
    task?: Task
}

function TaskForm({ onSubmit, task }: TaskFormProps) {
    const [dueDate, setDueDate] = useState(task ? task.due_date : undefined)
    const [title, setTitle] = useState(task ? task.title : '')
    const [description, setDescription] = useState(task ? task.description : '')
    const [subjectId, setSubjectId] = useState(task?.subject_id ? task.subject_id : '')

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (!dueDate) {
            alert('Empty due date')
            return
        }
        if (!title) {
            alert('Empty title.')
            return
        }

        try {
            if (onSubmit) {
                onSubmit({
                    dueDate: new Date(dueDate),
                    title: title,
                    description: description,
                    subjectId: subjectId !== '' ? subjectId : undefined,
                })
            }
        } catch (error) {
            alert('Failed to submit task.')
        }
    }

    return (
        <form onSubmit={handleSubmit}>
            <label htmlFor="due_date" className={styles.formLabel}>Due date</label>
            <br />
            <input
                type="date"
                name="due_date"
                className={styles.input}
                value={dueDate}
                onChange={e => setDueDate(e.target.value)}
            />
            <br />

            <label htmlFor="title" className={styles.formLabel}>Title</label>
            <br />
            <input
                type="text"
                name="title"
                className={styles.input}
                value={title}
                onChange={e => setTitle(e.target.value)}
            />
            <br />

            <label htmlFor="description" className={styles.formLabel}>Description</label>
            <br />
            <input
                type="text"
                name="description"
                className={styles.input}
                value={description}
                onChange={e => setDescription(e.target.value)}
            />
            <br />

            <label htmlFor="subject" className={styles.formLabel}>Subject</label>
            <br />
            <select
                name="subject"
                className={styles.input}
                value={subjectId}
                onChange={e => setSubjectId(e.target.value)}
            >
                <option value=""></option>
                {subjects.map(subject => (
                    <option key={subject['id']} value={subject['id']}>{subject['title']}</option>
                ))}
            </select>
            <br />

            <button type="submit" className={styles.addButton}>Submit</button>
        </form>
    )
}

export default TaskForm
