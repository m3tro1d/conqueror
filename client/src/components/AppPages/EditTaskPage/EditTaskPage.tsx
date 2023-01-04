import React from 'react'
import { useParams } from 'react-router-dom'

function EditTaskPage() {
    const { id } = useParams()

    return (
        <div>
            Edit task {id}
        </div>
    )
}

export default EditTaskPage
