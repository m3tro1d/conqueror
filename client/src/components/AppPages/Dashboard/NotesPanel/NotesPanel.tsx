import React from 'react'
import styles from './NotesPanel.module.css'
import Note from './Note/Note'

function NotesPanel() {
    const notes = [
        {
            id: '1',
            title: 'Test note',
            content: 'shes a dump shes a fucking nightmare',
            tags: [
                {
                    id: 't1',
                    name: 'ood',
                },
                {
                    id: 't2',
                    name: 'study',
                },
            ],
            updated_at: 'today at 15:00',
            subject_id: null,
        },
        {
            id: '2',
            title: 'Shut your fucking face, uncle fucker',
            content: 'got a little heartache hes a fucking weasel',
            tags: [],
            updated_at: 'yesterday',
            subject_id: 's1',
        },
    ]

    return (
        <div className={styles.notesPanel}>
            <ul className={styles.notesList}>
                {notes.map(note => (
                    <Note key={note.id} note={note} />
                ))}
            </ul>
        </div>
    )
}

export default NotesPanel
