import React, { FormEvent, useState } from 'react'
import { tasksApi } from '../../../../../api/api'

type AddTaskFormProps = {
    updateTasks: () => void
}

function AddTaskForm({ updateTasks }: AddTaskFormProps) {
    const [dueDate, setDueDate] = useState('')
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (title === '') {
            alert('Empty title.')
            return
        }

        try {
            await tasksApi.createTask({
                due_date: new Date(),
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
            <label htmlFor="due_date">Due date</label>
            <input
                type="date"
                name="due_date"
                onChange={e => setDueDate(e.target.value)}
            />
            <br />

            <label htmlFor="title">Title</label>
            <input
                type="text"
                name="title"
                onChange={e => setTitle(e.target.value)}
            />
            <br />

            <label htmlFor="description">Description</label>
            <input
                type="text"
                name="description"
                onChange={e => setDescription(e.target.value)}
            />
            <br />

            <button type="submit">Add</button>
        </form>
    )
}

export default AddTaskForm
