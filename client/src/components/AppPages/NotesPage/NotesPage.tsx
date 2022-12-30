import React from 'react'
import NoteForm from './AddNoteForm/NoteForm'
import useNotes from '../../../hooks/useNotes'
import Note from '../common/Note/Note'
import styles from './NotesPage.module.css'

function NotesPage() {
    const { notes, updateNotes } = useNotes()

    return (
        <div className={styles.notesPage}>
            <NoteForm updateNotes={updateNotes} />

            <ul className={styles.notesList}>
                {
                    notes.length === 0 &&
                    <div className={styles.noNotesNotice}>No notes</div>
                }
                {notes.map(note => (
                    <Note
                        key={note['id']}
                        note={note}
                    />
                ))}
            </ul>
        </div>
    )
}

export default NotesPage
