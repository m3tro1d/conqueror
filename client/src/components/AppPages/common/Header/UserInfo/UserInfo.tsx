import React, { useEffect, useRef, useState } from 'react'
import { userApi } from '../../../../../api/api'
import styles from './UserInfo.module.css'

function UserInfo(): JSX.Element {
    const [user, setUser] = useState(null)

    const updateUser = () => {
        userApi
            .getUser()
            .then(response => setUser(response))
            .catch(() => alert('Failed to fetch user info'))
    }


    const inputRef = useRef<HTMLInputElement | null>(null)
    const onChangeAvatar = (e: React.MouseEvent) => {
        e.preventDefault()
        inputRef?.current?.click()
    }
    const onSelectAvatar = (e: React.ChangeEvent<HTMLInputElement>) => {
        e.preventDefault()
        if (!e.target.files || e.target.files.length === 0) {
            alert('No file selected')
            return
        }

        userApi
            .changeAvatar(e.target.files[0])
            .then(updateUser)
            .catch(() => alert('Failed to set avatar'))
    }

    const logout = () => {
        localStorage.clear()
        window.location.assign('/')
    }

    useEffect(updateUser, [])

    return (
        <div className={styles.userInfo}>
            {
                user &&
                <div className={styles.userBlock}>
                    <span>{user['login']}</span>
                    {
                        user['avatar']
                            ? <img src={user['avatar']['url']} className={styles.avatar} onClick={onChangeAvatar} />
                            : <span onClick={onChangeAvatar} className={styles.avatarPlaceholder}>No image</span>
                    }
                    <a
                        href="#"
                        className={styles.logoutButton}
                        onClick={logout}
                    >
                        <span className="material-icons">logout</span>
                    </a>
                </div>
            }
            <input
                type="file"
                hidden
                accept="image/jpeg"
                ref={inputRef}
                onChange={onSelectAvatar}
            />
        </div>
    )
}

export default UserInfo
