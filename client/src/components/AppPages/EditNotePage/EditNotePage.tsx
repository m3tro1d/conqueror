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
        if (!id) {
            alert('Invalid note id')
            return
        }

        notesApi
            .updateNote(id, {
                title: note.title,
                content: note.content,
                subject_id: note.subjectId,
            })
            .then(() => window.location.assign('/notes'))
            .catch(() => alert('Failed to update note.'))
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
