import React, { useEffect, useState } from 'react'
import { authApi } from '../../../../../api/api'
import styles from './UserInfo.module.css'

function UserInfo(): JSX.Element {
    const [user, setUser] = useState({
        user_id: '',
        login: '',
    })
    useEffect(() => {
        authApi
            .getUser()
            .then(response => setUser(response))
            .catch(() => alert('Failed to fetch user info'))
    }, [])

    // TODO: profile photo!!!
    return (
        <div className={styles.userInfo}>
            <span>{user.login}</span>
        </div>
    )
}

export default UserInfo
