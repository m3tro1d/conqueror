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
    updated_at: number
    subject_id: string | null
    subject_title: string | null
}

type NoteProps = {
    note: Note
    removeNote?: (id: string) => void
}

function Note({ note, removeNote }: NoteProps) {
    const formatDate = (timestamp: number) => {
        const date = new Date(timestamp * 1000)
        return date.toLocaleString()
    }

    const trim = (text: string) => {
        if (text.length > 30) {
            return text.substring(0, 30) + '...'
        }

        return text
    }

    return (
        <li className={styles.noteItem}>
            <div className={styles.mainContent}>
                <span className={styles.updatedAt}>{formatDate(note.updated_at)}</span>
                <span className={styles.title}>
                    <a href={`/note/${note.id}`} className={styles.noteLink}>{trim(note.title)}</a>
                </span>
                <span className={styles.content}>{trim(note.content)}</span>
            </div>

            {
                note.subject_title !== null &&
                <span className={styles.subject}>
                    {note.subject_title}
                </span>
            }

            {
                removeNote &&
                <div>
                    <span
                        className={'material-icons ' + styles.removeButton}
                        onClick={e => {
                            e.preventDefault()
                            removeNote(note.id)
                        }}
                    >
                        delete
                    </span>
                </div>
            }
        </li>
    )
}

export default Note
