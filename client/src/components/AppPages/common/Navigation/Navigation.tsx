import React from 'react'
import styles from './Navigation.module.css'

function Navigation() {
    return (
        <div className={styles.navigation}>
            <ul className={styles.linksList}>
                <li className={styles.listItem}>
                    <a
                        href="/"
                        title="Dashboard"
                        className={styles.link}
                    >
                        <span className={"material-icons " + styles.icon}>home</span>
                    </a>
                </li>
                <li className={styles.listItem}>
                    <a
                        href="/tasks"
                        title="Tasks"
                        className={styles.link}
                    >
                        <span className={"material-icons " + styles.icon}>task</span>
                    </a>
                </li>
                <li className={styles.listItem}>
                    <a
                        href="/notes"
                        title="Notes"
                        className={styles.link}
                    >
                        <span className={"material-icons " + styles.icon}>edit_note</span>
                    </a>
                </li>
            </ul>
        </div>
    )
}

export default Navigation
