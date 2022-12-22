import React from 'react'
import styles from './Navigation.module.css'

function Navigation() {
    return (
        <div className={styles.navigation}>
            <ul className={styles.linksList}>
                <li><a href="/">Dashboard</a></li>
                <li><a href="/tasks">Tasks</a></li>
                <li><a href="/notes">Notes</a></li>
            </ul>
        </div>
    )
}

export default Navigation
