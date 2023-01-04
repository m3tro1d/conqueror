import React, {useEffect, useState} from 'react'
import NoteForm from '../common/NoteForm/NoteForm'
import useNotes from '../../../hooks/useNotes'
import Note from '../common/Note/Note'
import styles from './NotesPage.module.css'
import useDebounce from '../../../hooks/useDebounce'

function NotesPage() {
    const {notes, updateNotes, removeNote} = useNotes()

    const [query, setQuery] = useState('')
    const debouncedQuery = useDebounce(query, 500)
    useEffect(() => updateNotes(debouncedQuery), [debouncedQuery])

    return (
        <div className={styles.notesPage}>
            <NoteForm updateNotes={updateNotes}/>

            <label htmlFor="search" className={styles.searchLabel}>Search</label>
            <input
                type="text"
                name="search"
                className={styles.searchBar}
                onChange={e => setQuery(e.target.value)}
            />

            <ul className={styles.notesList}>
                {
                    notes.length === 0 &&
                    <div className={styles.noNotesNotice}>No notes</div>
                }
                {notes.map(note => (
                    <Note
                        key={note['id']}
                        note={note}
                        removeNote={removeNote}
                    />
                ))}
            </ul>
        </div>
    )
}

export default NotesPage
