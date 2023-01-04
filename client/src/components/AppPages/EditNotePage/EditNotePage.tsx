import React from 'react'
import { useParams } from 'react-router-dom'

function EditNotePage() {
    const { id } = useParams()

    return (
        <div>
            Edit note {id}
        </div>
    )
}

export default EditNotePage
