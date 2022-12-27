import React from 'react'
import styles from './Navigation.module.css'
import NavigationItem from './NavigationItem/NavigationItem'

function Navigation() {
    const links = [
        {
            href: '/',
            title: 'Dashboard',
            icon: 'home',
        },
        {
            href: '/subjects',
            title: 'Subjects',
            icon: 'width_normal',
        },
        {
            href: '/tasks',
            title: 'Tasks',
            icon: 'task',
        },
        {
            href: '/notes',
            title: 'Notes',
            icon: 'edit_note',
        },
    ]

    return (
        <div className={styles.navigation}>
            <ul className={styles.linksList}>
                {links.map(link => (
                    <NavigationItem
                        key={link.href}
                        href={link.href}
                        title={link.title}
                        icon={link.icon}
                    />
                ))}
            </ul>
        </div>
    )
}

export default Navigation
