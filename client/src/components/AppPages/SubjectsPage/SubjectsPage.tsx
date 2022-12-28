import React from 'react'
import styles from './SubjectsPage.module.css'
import Subject from './Subject/Subject'
import AddSubjectForm from './AddSubjectForm/AddSubjectForm'
import useSubjects from '../../../hooks/useSubjects'

function SubjectsPage() {
    const { subjects, updateSubjects, changeSubjectTitle, removeSubject } = useSubjects()

    return (
        <div className={styles.subjectsContainer}>
            <AddSubjectForm updateSubjects={updateSubjects} />
            <ul className={styles.subjectsList}>
                {subjects.map(subject => (
                    <Subject
                        key={subject['id']}
                        subject={subject}
                        changeSubjectTitle={changeSubjectTitle}
                        removeSubject={removeSubject}
                    />
                ))}
            </ul>
        </div>
    )
}

export default SubjectsPage
