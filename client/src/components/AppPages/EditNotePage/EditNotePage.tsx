import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { notesApi } from '../../../api/api'
import NoteForm, { NoteData } from '../common/NoteForm/NoteForm'

function EditNotePage() {
    const { id } = useParams()
    const [note, setNote] = useState()

    useEffect(() => {
        if (!id) {
            alert('Invalid note id')
            return
        }

        notesApi
            .getNote(id)
            .then(response => setNote(response.note))
            .catch(() => alert('Failed to fetch note'))
    }, [])

    const onSubmit = async (note: NoteData) => {
        // TODO: update note
        window.location.assign('/notes')
    }

    return (
        <div>
            <NoteForm
                onSubmit={onSubmit}
                note={note}
            />
        </div>
    )
}

export default EditNotePage
