import React from 'react'
import styles from './Note.module.css'

type Tag = {
    id: string
    name: string
}

type Note = {
    id: string
    title: string
    content: string
    tags: Tag[]
    updated_at: string
    subject_id: string | null
}

type NoteProps = {
    note: Note
}

function Note({ note }: NoteProps) {
    const trim = (text: string) => {
        if (text.length > 30) {
            return text.substring(0, 30) + '...'
        }

        return text
    }

    return (
        <li className={styles.noteItem}>
            <span className={styles.updatedAt}>{note.updated_at}</span>
            <span className={styles.title}>{trim(note.title)}</span>
            <span className={styles.content}>{trim(note.content)}</span>
        </li>
    )
}

export default Note
