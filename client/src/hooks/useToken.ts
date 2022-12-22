import { useState } from 'react'

function useToken() {
    const getToken = () => {
        const token = localStorage.getItem('token')
        return token ?? ''
    }

    const [token, setToken] = useState(getToken())

    const saveToken = (tokenString: string) => {
        localStorage.setItem('token', tokenString)
        setToken(tokenString)
    }

    return {
        token,
        setToken: saveToken,
    }
}

export default useToken
