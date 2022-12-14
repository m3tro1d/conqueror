import React from 'react'
import styles from './NotesPanel.module.css'
import useNotes from '../../../../hooks/useNotes'
import Note from '../../common/Note/Note'

function NotesPanel() {
    const { notes } = useNotes()

    return (
        <div className={styles.notesPanel}>
            <ul className={styles.notesList}>
                {
                    notes.length === 0 &&
                    <div className={styles.noNotesNotice}>No notes</div>
                }
                {notes.map(note => (
                    <Note key={note['id']} note={note}/>
                ))}
            </ul>
        </div>
    )
}

export default NotesPanel
