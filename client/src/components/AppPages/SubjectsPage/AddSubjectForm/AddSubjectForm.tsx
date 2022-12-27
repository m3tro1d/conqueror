import React, { FormEvent, useState } from 'react'
import { subjectApi } from '../../../../api/api'

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
            onSubmit={handleSubmit}
        >
            <input
                type="text"
                onChange={e => setTitle(e.target.value)}
            />

            <button type="submit">Add</button>
        </form>
    )
}

export default AddSubjectForm
