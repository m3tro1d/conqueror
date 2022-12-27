import React, { useEffect, useState } from 'react'
import styles from './SubjectsPage.module.css'
import { subjectApi } from '../../../api/api'
import Subject from './Subject/Subject'
import AddSubjectForm from './AddSubjectForm/AddSubjectForm'

function SubjectsPage() {
    const [subjects, setSubjects] = useState([])

    const updateSubjects = () => {
        subjectApi
            .listSubjects()
            .then(response => setSubjects(response.subjects))
            .catch(() => alert('Failed to fetch subjects.'))
    }
    const removeSubject = (id: string) => {
        subjectApi
            .removeSubject(id)
            .then(updateSubjects)
            .catch(() => alert('Failed to remove subject.'))
    }

    useEffect(updateSubjects, [])

    return (
        <div className={styles.subjectsContainer}>
            <AddSubjectForm updateSubjects={updateSubjects} />
            <ul className={styles.subjectsList}>
                {subjects.map(subject => (
                    <Subject
                        key={subject['id']}
                        subject={subject}
                        removeSubject={removeSubject}
                    />
                ))}
            </ul>
        </div>
    )
}

export default SubjectsPage
