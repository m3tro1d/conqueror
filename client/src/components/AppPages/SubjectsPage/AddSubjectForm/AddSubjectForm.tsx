import React, { FormEvent, useState } from 'react'
import { subjectApi } from '../../../../api/api'
import styles from './AddSubjectForm.module.css'

type AddSubjectFormProps = {
    updateSubjects: () => void
}

function AddSubjectForm({ updateSubjects }: AddSubjectFormProps) {
    const [title, setTitle] = useState('')

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        if (title === '') {
            alert('Empty title.')
            return
        }

        try {
            await subjectApi.createSubject({ title })
            updateSubjects()
        } catch (error) {
            alert('Failed to add subject.')
        }
    }

    return (
        <form
            className={styles.form}
            onSubmit={handleSubmit}
        >
            <input
                type="text"
                className={styles.input}
                onChange={e => setTitle(e.target.value)}
            />

            <button
                type="submit"
                className={styles.submitButton}
            >
                Add
            </button>
        </form>
    )
}

export default AddSubjectForm
