import React, { FormEvent, useRef } from 'react'
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
    const titleRef = useRef<HTMLInputElement | null>(null)
    const contentRef = useRef<HTMLTextAreaElement | null>(null)
    const subjectIdRef = useRef<HTMLSelectElement | null>(null)

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (!titleRef.current?.value) {
            alert('Empty title')
            return
        }
        if (!contentRef.current) {
            return
        }

        try {
            if (onSubmit) {
                onSubmit({
                    title: titleRef.current.value,
                    content: contentRef.current.value,
                    subjectId: subjectIdRef.current?.value ? subjectIdRef.current.value : undefined,
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
                defaultValue={note ? note.title : ''}
                ref={titleRef}
            />
            <br />

            <label htmlFor="content" className={styles.formLabel}>Content</label>
            <br />
            <textarea
                name="content"
                className={styles.content}
                defaultValue={note ? note.content : ''}
                ref={contentRef}
            ></textarea>
            <br />

            <label htmlFor="subject" className={styles.formLabel}>Subject</label>
            <br />
            <select
                name="subject"
                className={styles.input}
                defaultValue={note?.subject_id ? note.subject_id : ''}
                ref={subjectIdRef}
            >
                <option value=""></option>
                {subjects.map(subject => (
                    <option
                        key={subject['id']}
                        value={subject['id']}
                        selected={note?.subject_id === subject['id']}
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

export default NoteForm
