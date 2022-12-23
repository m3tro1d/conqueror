import React from 'react'
import styles from './NavigationItem.module.css'

type NavigationItemProps = {
    href: string
    title: string
    icon: string
}

function NavigationItem({ href, title, icon }: NavigationItemProps) {
    const getClassName = (href: string) => {
        if (document.location.pathname == href) {
            return styles.activeLink
        }
        return styles.link
    }

    return (
        <li className={styles.listItem}>
            <a
                href={href}
                title={title}
                className={getClassName(href)}
            >
                <span className={'material-icons ' + styles.icon}>
                    {icon}
                </span>
            </a>
        </li>
    )
}

export default NavigationItem
