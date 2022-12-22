import React, { useEffect, useState } from 'react'
import { authApi } from '../../../../../api/api'

function UserInfo(): JSX.Element {
    const [user, setUser] = useState({
        user_id: '',
        login: '',
    })
    useEffect(() => {
        authApi
            .getUser()
            .then(response => setUser(response))
            .catch(error => alert('Failed to fetch user info'))
    }, [])

    return (
        <div>
            <span>{user.login}</span>
            <div>Profile photo</div>
        </div>
    )
}

export default UserInfo
