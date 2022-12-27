import React, { FormEvent, useState } from 'react'
import styles from './LoginPage.module.css'
import { authApi } from '../../../api/api'

type LoginPageProps = {
    setToken: (token: string) => void
}

function LoginPage({ setToken }: LoginPageProps) {
    const [login, setLogin] = useState('')
    const [password, setPassword] = useState('')

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        try {
            const response = await authApi.login({
                login,
                password,
            })
            setToken(response.token)
            document.location.assign('/')
        } catch (error) {
            alert('Failed to login. Please check username and password.')
        }
    }

    return (
        <form
            className={styles.form}
            onSubmit={handleSubmit}
        >
            <h1 className={styles.header}>Conqueror</h1>

            <input
                type="text"
                name="login"
                placeholder="Login"
                className={styles.input}
                onChange={e => setLogin(e.target.value)}
            />
            <br />

            <input
                type="password"
                name="password"
                placeholder="Password"
                className={styles.input}
                onChange={e => setPassword(e.target.value)}
            />
            <br />

            <div className={styles.buttonsContainer}>
                <button type="submit" className={styles.loginButton}>Login</button>
                <button
                    className={styles.signupButton}
                    onClick={e => {
                        e.preventDefault()
                        document.location.assign('/signup')
                    }}
                >
                    Sign up
                </button>
            </div>
        </form>
    )
}

export default LoginPage
