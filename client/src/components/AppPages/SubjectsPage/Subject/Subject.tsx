import React from 'react'
import styles from './Subject.module.css'

type Subject = {
    id: string
    title: string
}

type SubjectProps = {
    subject: Subject
    removeSubject: (id: string) => void
}

function Subject({ subject, removeSubject }: SubjectProps) {
    return (
        <li
            key={subject.id}
            className={styles.listItem}
        >
            <span className={styles.subjectTitle}>{subject.title}</span>
            <a
                href="#"
                className={styles.removeButton}
                onClick={e => {
                    e.preventDefault()
                    removeSubject(subject.id)
                }}
            >
                <span className="material-icons">delete</span>
            </a>
        </li>
    )
}

export default Subject
