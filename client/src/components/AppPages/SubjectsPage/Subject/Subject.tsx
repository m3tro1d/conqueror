import React from 'react'
import styles from './Subject.module.css'

type Subject = {
    id: string
    title: string
}

type SubjectProps = {
    subject: Subject
    changeSubjectTitle: (id: string, title: string) => void
    removeSubject: (id: string) => void
}

function Subject({ subject, changeSubjectTitle, removeSubject }: SubjectProps) {
    const onChange = (e: React.MouseEvent) => {
        e.preventDefault()
        const title = prompt('Enter new title', subject.title)
        if (title === null || title === '') {
            alert('Empty title.')
            return
        }
        changeSubjectTitle(subject.id, title)
    }
    const onRemove = (e: React.MouseEvent) => {
        e.preventDefault()
        removeSubject(subject.id)
    }

    return (
        <li
            key={subject.id}
            className={styles.listItem}
        >
            <span className={styles.subjectTitle}>{subject.title}</span>
            <a
                href="#"
                className={styles.editButton}
                onClick={onChange}
            >
                <span className="material-icons">edit</span>
            </a>
            <a
                href="#"
                className={styles.removeButton}
                onClick={onRemove}
            >
                <span className="material-icons">delete</span>
            </a>
        </li>
    )
}

export default Subject
