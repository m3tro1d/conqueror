import { useEffect, useState } from 'react'
import { subjectApi } from '../api/api'

function useSubjects() {
    const [subjects, setSubjects] = useState([])

    const updateSubjects = () => {
        subjectApi
            .listSubjects()
            .then(response => setSubjects(response.subjects))
            .catch(() => alert('Failed to fetch subjects.'))
    }
    const changeSubjectTitle = (id: string, title: string) => {
        subjectApi
            .changeSubjectTitle(id, title)
            .then(updateSubjects)
            .catch(() => alert('Failed to change subject title.'))
    }
    const removeSubject = (id: string) => {
        subjectApi
            .removeSubject(id)
            .then(updateSubjects)
            .catch(() => alert('Failed to remove subject.'))
    }

    useEffect(updateSubjects, [])

    return {
        subjects,
        updateSubjects,
        changeSubjectTitle,
        removeSubject,
    }
}

export default useSubjects
