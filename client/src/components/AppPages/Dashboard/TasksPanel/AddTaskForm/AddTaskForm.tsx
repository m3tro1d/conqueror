import React, { FormEvent, useState } from 'react'
import { tasksApi } from '../../../../../api/api'
import styles from './AddTaskForm.module.css'

type AddTaskFormProps = {
    updateTasks: () => void
}

function AddTaskForm({ updateTasks }: AddTaskFormProps) {
    const [dueDate, setDueDate] = useState<Date | null>(null)
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')

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

            <button type="submit" className={styles.addButton}>Add</button>
        </form>
    )
}

export default AddTaskForm
