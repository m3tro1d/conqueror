import React, { FormEvent, useState } from 'react'
import { tasksApi } from '../../../../../api/api'
import styles from './AddTaskForm.module.css'
import useSubjects from '../../../../../hooks/useSubjects'

type AddTaskFormProps = {
    updateTasks: () => void
}

function AddTaskForm({ updateTasks }: AddTaskFormProps) {
    const [dueDate, setDueDate] = useState<Date | null>(null)
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    const [subjectId, setSubjectId] = useState('')

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (dueDate === null) {
            alert('Empty due date')
            return
        }
        if (title === '') {
            alert('Empty title.')
            return
        }

        try {
            await tasksApi.createTask({
                due_date: dueDate,
                title: title,
                description: description,
                subject_id: subjectId !== '' ? subjectId : undefined,
            })
            updateTasks()
        } catch (error) {
            alert('Failed to add task.')
        }
    }

    return (
        <form
            onSubmit={handleSubmit}
        >
            <label htmlFor="due_date" className={styles.formLabel}>Due date</label>
            <br />
            <input
                type="date"
                name="due_date"
                className={styles.input}
                onChange={e => setDueDate(e.target.valueAsDate)}
            />
            <br />

            <label htmlFor="title" className={styles.formLabel}>Title</label>
            <br />
            <input
                type="text"
                name="title"
                className={styles.input}
                onChange={e => setTitle(e.target.value)}
            />
            <br />

            <label htmlFor="description" className={styles.formLabel}>Description</label>
            <br />
            <input
                type="text"
                name="description"
                className={styles.input}
                onChange={e => setDescription(e.target.value)}
            />
            <br />

            <label htmlFor="subject" className={styles.formLabel}>Subject</label>
            <br />
            <select
                name="subject"
                className={styles.input}
                onChange={e => setSubjectId(e.target.value)}
            >
                <option value=""></option>
                {subjects.map(subject => (
                    <option key={subject['id']} value={subject['id']}>{subject['title']}</option>
                ))}
            </select>
            <br />

            <button type="submit" className={styles.addButton}>Add</button>
        </form>
    )
}

export default AddTaskForm
