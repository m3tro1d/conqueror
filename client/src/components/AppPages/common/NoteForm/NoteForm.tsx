import React, { FormEvent, useState } from 'react'
import useSubjects from '../../../../hooks/useSubjects'
import { notesApi } from '../../../../api/api'
import styles from './NoteForm.module.css'

type NoteFormProps = {
    updateNotes: () => void
}

function NoteForm({ updateNotes }: NoteFormProps) {
    const [title, setTitle] = useState('')
    const [content, setContent] = useState('')
    const [subjectId, setSubjectId] = useState('')

    const { subjects } = useSubjects()

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (title === '') {
            alert('Empty title')
            return
        }

        try {
            await notesApi.createNote({
                title: title,
                content: content,
                subject_id: subjectId !== '' ? subjectId : undefined,
            })
            updateNotes()
        } catch (error) {
            alert('Failed to add note.')
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
                onChange={e => setTitle(e.target.value)}
            />
            <br />

            <label htmlFor="content" className={styles.formLabel}>Content</label>
            <br />
            <textarea
                name="content"
                className={styles.content}
                onChange={e => setContent(e.target.value)}
            ></textarea>
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

export default NoteForm
