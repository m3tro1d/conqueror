import React, { FormEvent, useRef } from 'react'
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
    const dueDateRef = useRef<HTMLInputElement | null>(null)
    const titleRef = useRef<HTMLInputElement | null>(null)
    const descriptionRef = useRef<HTMLInputElement | null>(null)
    const subjectIdRef = useRef<HTMLSelectElement | null>(null)

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (!dueDateRef.current?.value) {
            alert('Empty due date')
            return
        }
        if (!titleRef.current?.value) {
            alert('Empty title.')
            return
        }
        if (!descriptionRef.current) {
            return
        }

        try {
            if (onSubmit) {
                onSubmit({
                    dueDate: new Date(dueDateRef.current.value),
                    title: titleRef.current.value,
                    description: descriptionRef.current.value,
                    subjectId: subjectIdRef.current?.value ? subjectIdRef.current.value : undefined,
                })
            }
        } catch (error) {
            alert('Failed to submit task.')
        }
    }

    const dateDefaultValue = (() => {
        if (task) {
            const date = new Date(Date.parse(task.due_date))
            const month = (date.getMonth() + 1).toString().padStart(2, '0')
            const day = date.getDate().toString().padStart(2, '0')

            return `${date.getFullYear()}-${month}-${day}`
        }

        return ''
    })()

    return (
        <form onSubmit={handleSubmit}>
            <label htmlFor="due_date" className={styles.formLabel}>Due date</label>
            <br />
            <input
                type="date"
                name="due_date"
                className={styles.input}
                defaultValue={dateDefaultValue}
                ref={dueDateRef}
            />
            <br />

            <label htmlFor="title" className={styles.formLabel}>Title</label>
            <br />
            <input
                type="text"
                name="title"
                className={styles.input}
                defaultValue={task ? task.title : ''}
                ref={titleRef}
            />
            <br />

            <label htmlFor="description" className={styles.formLabel}>Description</label>
            <br />
            <input
                type="text"
                name="description"
                className={styles.input}
                defaultValue={task ? task.description : ''}
                ref={descriptionRef}
            />
            <br />

            <label htmlFor="subject" className={styles.formLabel}>Subject</label>
            <br />
            <select
                name="subject"
                className={styles.input}
                ref={subjectIdRef}
            >
                <option value=""></option>
                {subjects.map(subject => (
                    <option
                        key={subject['id']}
                        value={subject['id']}
                        selected={task?.subject_id === subject['id']}
                    >
                        {subject['title']}
                    </option>
                ))}
            </select>
            <br />

            <button type="submit" className={styles.addButton}>Submit</button>
        </form>
    )
}

export default TaskForm
