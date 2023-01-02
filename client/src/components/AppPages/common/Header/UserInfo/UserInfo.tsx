import React, {useEffect, useRef, useState} from 'react'
import {userApi} from '../../../../../api/api'
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

    useEffect(updateUser, [])

    return (
        <div className={styles.userInfo}>
            {
                user &&
                <>
                    <span>{user['login']}</span>
                    {
                        user['avatar']
                            ? <img src={user['avatar']['url']} className={styles.avatar} onClick={onChangeAvatar}/>
                            : <span onClick={onChangeAvatar}>No image</span>
                    }
                </>
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
