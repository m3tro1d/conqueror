import { useEffect, useState } from 'react'
import { notesApi } from '../api/api'

function useNotes() {
    const [notes, setNotes] = useState([])

    const updateNotes = () => {
        notesApi
            .listNotes()
            .then(response => setNotes(response.notes))
            .catch(() => alert('Failed to fetch notes.'))
    }
    const removeNote = (id: string) => {
        notesApi
            .removeNote(id)
            .then(updateNotes)
            .catch(() => alert('Failed to remove note.'))
    }

    useEffect(updateNotes, [])

    return {
        notes,
        updateNotes,
        removeNote,
    }
}

export default useNotes
