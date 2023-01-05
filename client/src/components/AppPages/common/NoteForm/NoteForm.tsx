import React, { FormEvent, useState } from 'react'
import useSubjects from '../../../../hooks/useSubjects'
import styles from './NoteForm.module.css'

type Note = {
    id: string
    title: string
    content: string
    updated_at: number
    subject_id: string | null
    subject_title: string | null
}

export type NoteData = {
    title: string
    content: string
    subjectId?: string
}

type NoteFormProps = {
    onSubmit: (note: NoteData) => void
    note?: Note
}

function NoteForm({ onSubmit, note }: NoteFormProps) {
    const [title, setTitle] = useState(note ? note.title : '')
    const [content, setContent] = useState(note ? note.content : '')
    const [subjectId, setSubjectId] = useState(note?.subject_id ? note.subject_id : '')

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (title === '') {
            alert('Empty title')
            return
        }

        try {
            if (onSubmit) {
                onSubmit({
                    title: title,
                    content: content,
                    subjectId: subjectId !== '' ? subjectId : undefined,
                })
            }
        } catch (error) {
            alert('Failed to submit note.')
        }
    }

    return (
        <form onSubmit={handleSubmit}>
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

            <label htmlFor="content" className={styles.formLabel}>Content</label>
            <br />
            <textarea
                name="content"
                className={styles.content}
                value={content}
                onChange={e => setContent(e.target.value)}
            ></textarea>
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

export default NoteForm
