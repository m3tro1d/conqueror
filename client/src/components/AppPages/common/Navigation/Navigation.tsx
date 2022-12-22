import React from 'react'
import styles from './Navigation.module.css'

function Navigation() {
    return (
        <div className={styles.navigation}>
            <ul className={styles.linksList}>
                <li>
                    <a href="/" title="Dashboard">
                        <span className="material-icons">home</span>
                    </a>
                </li>
                <li>
                    <a href="/tasks" title="Tasks">
                        <span className="material-icons">task</span>
                    </a>
                </li>
                <li>
                    <a href="/notes" title="Notes">
                        <span className="material-icons">edit_note</span>
                    </a>
                </li>
            </ul>
        </div>
    )
}

export default Navigation
